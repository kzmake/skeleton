package handler

import (
	"context"

	pb "github.com/kzmake/skeleton/gen/go/echo/v1"
)

type echo struct {
	pb.UnimplementedEchoServer
}

var _ pb.EchoServer = new(echo)

func NewEcho() pb.EchoServer { return &echo{} }

func (h *echo) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Msg: req.Msg}, nil
}
