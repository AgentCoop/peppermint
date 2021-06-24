package server

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/balancer"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/server"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	grpc2 "github.com/AgentCoop/peppermint/pkg/grpc"
	balancer2 "github.com/AgentCoop/peppermint/pkg/service/balancer"
	"google.golang.org/grpc"
	"net"
)

type HubServer interface {
	grpc2.BaseServer
	balancer.BalancerServer
}

type balancerServer struct {
	grpc2.BaseServer
	balancer2.Balancer
	hub.UnimplementedHubServer
}

func (s *balancerServer) unknownStreamHandler(srv interface{}, stream grpc.ServerStream) error {
	err := s.Balancer.ForwardCall(srv, stream)
	return err
}

func withUnaryServerMiddlewares(svcName string) grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		middleware.PreUnaryInterceptor(svcName),
		middleware.PostUnaryInterceptor(svcName),
	)
}

func NewServer(svcName string, address net.Addr, balancer balancer2.Balancer) *balancerServer {
	s := new(balancerServer)
	s.Balancer = balancer
	s.BaseServer = server.NewBaseServer(svcName, address, grpc.NewServer(
		grpc.UnknownServiceHandler(s.unknownStreamHandler),
		withUnaryServerMiddlewares(svcName),
	))
	s.RegisterServer()
	return s
}

func (h *balancerServer) RegisterServer() {
	hub.RegisterHubServer(h.Handle(), h)
}
