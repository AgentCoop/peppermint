package hub

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/server"
	md_middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/server/metadata"
	"github.com/AgentCoop/peppermint/internal/grpc/server"

	"google.golang.org/grpc"
)

type HubServer interface {
	server.BaseServer
	hub.HubServer
}

type hubServer struct {
	server.BaseServer
	hub.UnimplementedHubServer
}

func withUnaryServerMiddlewares() grpc.ServerOption {
	return middleware.WithUnaryServerChain(
		md_middleware.UnaryServerInterceptor(),
	)
}

func NewServer(address string) *hubServer {
	s := new(hubServer)
	s.BaseServer = server.NewBaseServer(address, grpc.NewServer(
		withUnaryServerMiddlewares(),
	))
	s.RegisterServer()
	return s
}

func (h *hubServer) RegisterServer() {
	hub.RegisterHubServer(h.Handle(), h)
}
