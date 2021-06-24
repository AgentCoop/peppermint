package server

import (
	api "github.com/AgentCoop/peppermint/internal/api/peppermint/service/frontoffice/test"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/server"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc"
	"net"
)

type testServer struct {
	runtime.BaseServer
	api.TestServer
	hub.UnimplementedHubServer
}

func withUnaryServerMiddlewares(svcName string) grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		middleware.PreUnaryInterceptor(svcName),
		middleware.PostUnaryInterceptor(svcName),
	)
}

func NewServer(svcName string, address net.Addr) *testServer {
	s := new(testServer)
	s.BaseServer = server.NewBaseServer(svcName, address, grpc.NewServer(
		withUnaryServerMiddlewares(svcName),
	))
	s.RegisterServer()
	return s
}

func (h *testServer) RegisterServer() {
	hub.RegisterHubServer(h.Handle(), h)
}
