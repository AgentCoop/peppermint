package server

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/server"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	grpc2 "github.com/AgentCoop/peppermint/pkg/grpc"
	"google.golang.org/grpc"
	"net"
)

type HubServer interface {
	grpc2.BaseServer
	hub.HubServer
}

type hubServer struct {
	grpc2.BaseServer
	hub.UnimplementedHubServer
}

func withUnaryServerMiddlewares(svcName string) grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		middleware.PreUnaryInterceptor(svcName),
		middleware.SessionUnaryInterceptor(svcName),
		middleware.PostUnaryInterceptor(svcName),
	)
}

func NewServer(svcName string, address net.Addr) *hubServer {
	s := new(hubServer)
	s.BaseServer = server.NewBaseServer(svcName, address, grpc.NewServer(
		withUnaryServerMiddlewares(svcName),
	))
	s.RegisterServer()
	return s
}

func (h *hubServer) RegisterServer() {
	//h.Handle().
	hub.RegisterHubServer(h.Handle(), h)
}
