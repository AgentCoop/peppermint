package server

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/server"
	md_middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/server/metadata"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"net"

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

func NewServer(name string, address net.Addr) *hubServer {
	s := new(hubServer)
	s.BaseServer = server.NewBaseServer(name, address, grpc.NewServer(
		withUnaryServerMiddlewares(),
	))
	s.RegisterServer()
	return s
}

func (h *hubServer) RegisterServer() {
	hub.RegisterHubServer(h.Handle(), h)
}
