package server

import (
	_ "github.com/AgentCoop/peppermint/internal/grpc/codec"
	"google.golang.org/grpc"
	"net"
)

type baseServer struct {
	fullName string
	address  net.Addr
	handle   *grpc.Server
	lis      net.Listener
}

func (s *baseServer) Address() net.Addr {
	return s.address
}

func (s *baseServer) RegisterServer() {
	panic("implement me")
}

func NewBaseServer(fullName string, address net.Addr, server *grpc.Server) *baseServer {
	s := new(baseServer)
	s.fullName = fullName
	s.address = address
	s.handle = server
	return s
}

func (s *baseServer) Handle() *grpc.Server {
	return s.handle
}

func (s *baseServer) FullName() string {
	return s.fullName
}

func (s *baseServer) Methods() []string {
	info := s.handle.GetServiceInfo()
	_ = info
	return nil
}
