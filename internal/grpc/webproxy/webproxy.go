package webproxy

import (
	"crypto/tls"
	job "github.com/AgentCoop/go-work"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/server"
	md_middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/server/metadata"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

type webproxy struct {
	server.BaseServer
	lis net.Listener
	tlsHttpServer *http.Server
	x509CertPEM []byte
	x509KeyPEM []byte
	serverName string
}

func withUnaryServerMiddlewares() grpc.ServerOption {
	return middleware.WithUnaryServerChain(
		md_middleware.UnaryServerInterceptor(),
	)
}

func NewServer(address net.Addr, serverName string, x509CertPem []byte, x509KeyPem []byte) *webproxy {
	s := new(webproxy)
	s.BaseServer = server.NewBaseServer(address, grpc.NewServer(
		withUnaryServerMiddlewares(),
	))
	s.serverName = serverName
	s.x509CertPEM = x509CertPem
	s.x509KeyPEM = x509KeyPem
	s.RegisterServer()
	return s
}

func (b *webproxy) RegisterServer() {
	//t.RegisterTestServer(b.Handle(), b)
}

func (w *webproxy) StartTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		cer, err := tls.X509KeyPair(w.x509CertPEM, w.x509KeyPEM)
		task.Assert(err)
		tlsCfg := tls.Config{
			Certificates:                []tls.Certificate{cer},
			ServerName:                  w.serverName,
		}
		w.tlsHttpServer = &http.Server{
			Addr:              w.BaseServer.Address().String(),
			TLSConfig:         &tlsCfg,
		}
		wrappedGrpc := grpcweb.WrapServer(w.Handle())
		w.tlsHttpServer.Handler = http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			if wrappedGrpc.IsGrpcWebRequest(req) {
				wrappedGrpc.ServeHTTP(resp, req)
				return
			}
			// Fall back to other servers.
			http.DefaultServeMux.ServeHTTP(resp, req)
		})
		lis, err := tls.Listen("tcp", w.BaseServer.Address().String(), &tlsCfg)
		task.Assert(err)
		w.lis = lis
	}
	run := func(task job.Task) {
		w.tlsHttpServer.Serve(w.lis)
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}
