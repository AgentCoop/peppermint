package server

import (
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"github.com/AgentCoop/peppermint/internal/service"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

type webproxy struct {
	server.BaseServer
	service.WebProxyBalancer
	lis net.Listener
	tlsHttpServer *http.Server
	x509CertPEM []byte
	x509KeyPEM []byte
	serverName string
}

func withUnaryServerMiddlewares(serviceName string) grpc.ServerOption {
	//return middleware.WithUnaryServerChain(
	//	nil,
	//)
	return nil
}

func (b *webproxy) unknownStreamHandler(srv interface{}, stream grpc.ServerStream) error {
	err := b.WebProxyBalancer.ForwardCall(srv, stream)
	return err
}

func NewServer(name string, address net.Addr, cfg service.WebProxyConfigurator, balancer service.WebProxyBalancer) *webproxy {
	s := new(webproxy)
	s.BaseServer = server.NewBaseServer(name, address, grpc.NewServer(
		withUnaryServerMiddlewares(name),
		grpc.UnknownServiceHandler(s.unknownStreamHandler),
	))
	s.WebProxyBalancer = balancer
	s.serverName = cfg.ServerName()
	s.x509CertPEM = cfg.X509KeyPEM()
	s.x509KeyPEM = cfg.X509CertPEM()
	s.RegisterServer()
	return s
}

func (b *webproxy) RegisterServer() {
	//t.RegisterTestServer(b.Handle(), b)
}
