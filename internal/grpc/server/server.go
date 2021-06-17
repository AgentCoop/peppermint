package server

import (
	job "github.com/AgentCoop/go-work"
	_ "github.com/AgentCoop/peppermint/internal/grpc/codec"
	"google.golang.org/grpc"
	"net"
)

type baseServer struct {
	fullName string
	address  net.Addr
	handle   *grpc.Server
	lis      net.Listener
	logger   job.LogHandler
}

func (s *baseServer) Address() net.Addr {
	return s.address
}

func (s *baseServer) WithStdoutLogger(handler job.LogHandler) {
	s.logger = handler
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
	for name, serviceInfo := range s.handle.GetServiceInfo() {
		if name == s.fullName {
			out := make([]string, len(serviceInfo.Methods))
			for i := 0; i < len(serviceInfo.Methods); i++ {
				out[i] = serviceInfo.Methods[i].Name
			}
			return out
		}
	}
	return nil
}
