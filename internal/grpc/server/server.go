package server

import (
	_ "github.com/AgentCoop/peppermint/internal/grpc/codec"

	"google.golang.org/grpc"
)

type server struct {
	grpc *grpc.Server
}

func DefaultGrpcServer() *grpc.Server {
	opts := make([]grpc.ServerOption, 0)

	//var opts = []grpc.ServerOption{
	//	grpc.CallContentSubtype("dd"),
	//}
	return grpc.NewServer(opts...)
}

