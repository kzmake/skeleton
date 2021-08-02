package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	ginlogger "github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "github.com/kzmake/skeleton/gen/go/greeter/v1"
)

type Env struct {
	Address string `default:"0.0.0.0:8080"`
	Dapr    struct {
		Address string `default:"localhost:50051"`
	}
}

const prefix = "GATEWAY"

var env Env

func init() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

	if err := envconfig.Process(prefix, &env); err != nil {
		log.Fatal().Msgf("%+v", err)
	}

	log.Debug().Msgf("%+v", env)
}

func newGatewayServer(ctx context.Context) (*http.Server, error) {
	greeterMux := runtime.NewServeMux(
		runtime.WithMetadata(func(_ context.Context, _ *http.Request) metadata.MD { return metadata.Pairs("dapr-app-id", "greeter") }),
	)

	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := pb.RegisterGreeterHandlerFromEndpoint(ctx, greeterMux, env.Dapr.Address, opts); err != nil {
		return nil, xerrors.Errorf("Failed to register handler: %w", err)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(ginlogger.SetLogger(
		ginlogger.WithLogger(func(c *gin.Context, _ io.Writer, latency time.Duration) zerolog.Logger {
			return log.Logger.With().
				Timestamp().
				Int("status", c.Writer.Status()).
				Str("method", c.Request.Method).
				Str("path", c.Request.URL.Path).
				Str("ip", c.ClientIP()).
				Dur("latency", latency).
				Str("user_agent", c.Request.UserAgent()).
				Logger()
		}),
		ginlogger.WithSkipPath([]string{"/healthz"}),
	))
	r.Use(gin.Recovery())

	r.Any("/*any", gin.WrapH(greeterMux))

	return &http.Server{Addr: env.Address, Handler: r}, nil
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	gatewayS, err := newGatewayServer(ctx)
	if err != nil {
		log.Fatal().Msgf("Failed to build gateway server: %v", err)
	}
	g.Go(func() error { return gatewayS.ListenAndServe() })

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	select {
	case <-quit:
		break
	case <-ctx.Done():
		break
	}

	cancel()

	log.Info().Msg("Shutting down server...")

	ctx, timeout := context.WithTimeout(context.Background(), 5*time.Second)
	defer timeout()

	if err := gatewayS.Shutdown(ctx); err != nil {
		return xerrors.Errorf("failed to shutdown: %w", err)
	}

	log.Info().Msgf("Server exiting")

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal().Msgf("Failed to run server: %v", err)
	}
}
