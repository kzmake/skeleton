package handler

import (
	"context"

	pb "github.com/kzmake/skeleton/gen/go/greeter/v1"
)

type greeter struct {
	pb.UnimplementedGreeterServer
}

var _ pb.GreeterServer = new(greeter)

func NewGreeter() pb.GreeterServer { return &greeter{} }

func (h *greeter) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Msg: "Hello, " + req.GetName()}, nil
}
