package server

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/server"
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

func withUnaryServerMiddlewares(svcName string) grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		middleware.PreUnaryInterceptor(svcName),
		middleware.PostUnaryInterceptor(svcName),
	)
}

func NewServer(name string, address net.Addr) *hubServer {
	s := new(hubServer)
	s.BaseServer = server.NewBaseServer(name, address, grpc.NewServer(
		withUnaryServerMiddlewares(name),
	))
	s.RegisterServer()
	return s
}

func (h *hubServer) RegisterServer() {
	hub.RegisterHubServer(h.Handle(), h)
}
