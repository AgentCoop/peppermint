package hub

import (
	"google.golang.org/grpc"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/server"
)

func withUnaryServerMiddlewares() grpc.ServerOption {
	return middleware.WithUnaryServerChain()
}
