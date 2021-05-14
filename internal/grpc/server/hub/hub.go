package hub

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
)

type HubGrpcServer interface {
	server.GrpcServer
}

type hubGrpcServer struct {
	hub.UnimplementedHubServer
	server.BaseServer
}

func NewServer(address string) *hubGrpcServer {
	s := &hubGrpcServer{}
	s.Address = address
	return s
}

func (h *hubGrpcServer) RegisterService() {
	hub.RegisterHubServer(h.Handle, h)
}
