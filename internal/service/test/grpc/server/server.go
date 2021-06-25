package server

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	api "github.com/AgentCoop/peppermint/internal/api/peppermint/service/frontoffice/test"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/server"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	g "github.com/AgentCoop/peppermint/pkg/grpc"
	"google.golang.org/grpc"
	"net"
)

type testServer struct {
	g.BaseServer
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
