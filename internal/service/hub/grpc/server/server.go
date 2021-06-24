package server

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/server"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"net"
	"google.golang.org/grpc"
)

type HubServer interface {
	runtime.BaseServer
	hub.HubServer
}

type hubServer struct {
	runtime.BaseServer
	hub.UnimplementedHubServer
}

func withUnaryServerMiddlewares(svcName string) grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		middleware.PreUnaryInterceptor(svcName),
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
