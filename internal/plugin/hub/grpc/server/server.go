package server

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/server"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"github.com/AgentCoop/peppermint/internal/utils"
	"net"

	"google.golang.org/grpc"
)

type HubServer interface {
	g.BaseServer
	hub.HubServer
}

type hubServer struct {
	g.BaseServer
	hub.UnimplementedHubServer
}

func withUnaryServerMiddlewares(svcName string) grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		middleware.PreUnaryInterceptor(svcName),
		middleware.PostUnaryInterceptor(svcName),
	)
}

func NewServer(fullName string, address net.Addr) *hubServer {
	s := new(hubServer)
	name := utils.Grpc_ExtractServerShortName(fullName)
	s.BaseServer = server.NewBaseServer(fullName, address, grpc.NewServer(
		withUnaryServerMiddlewares(name),
	))
	s.RegisterServer()
	return s
}

func (h *hubServer) RegisterServer() {
	//h.Handle().
	hub.RegisterHubServer(h.Handle(), h)
}
