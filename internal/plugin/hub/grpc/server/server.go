package server

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	g "github.com/AgentCoop/peppermint/internal/grpc"
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

func encryptionLayer(desc g.CallDesc, encKey []byte) {
	//desc.WithEncKey(nil)
}

func encUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		callDesc := ctx.(g.ServerCallDesc)
		encryptionLayer(callDesc, nil)
		r, err := handler(ctx, req)
		return r, err
	}
}

func withUnaryServerMiddlewares(svcName string) grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		middleware.PreUnaryInterceptor(svcName),
		encUnaryInterceptor(),
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
