package webproxy

import (
	"crypto/tls"
	"fmt"
	job "github.com/AgentCoop/go-work"
	t "github.com/AgentCoop/peppermint/internal/api/peppermint/service/test"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/server"
	md_middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/server/metadata"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

type webproxy struct {
	server.BaseServer
	//t.TestServer
	t.UnimplementedTestServer

	lis net.Listener
	tlsHttpServer *http.Server
}

func withUnaryServerMiddlewares() grpc.ServerOption {
	return middleware.WithUnaryServerChain(
		md_middleware.UnaryServerInterceptor(),
	)
}

func NewServer(address string) *webproxy {
	s := new(webproxy)
	s.BaseServer = server.NewBaseServer(address, grpc.NewServer(
		withUnaryServerMiddlewares(),
	))
	s.RegisterServer()
	return s
}

func (b *webproxy) RegisterServer() {
	t.RegisterTestServer(b.Handle(), b)
}

func (w *webproxy) StartTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		certBaseDir  := "/home/pihpah/mycerts/peppermint.io/"
		cer, err := tls.LoadX509KeyPair(certBaseDir + "server.crt", certBaseDir + "server.key")
		task.Assert(err)

		tlsCfg := tls.Config{
			Certificates:                []tls.Certificate{cer},
			ServerName:                  "peppermint.io",
		}
		
		w.tlsHttpServer = &http.Server{
			Addr:              w.BaseServer.Address(),
			TLSConfig:         &tlsCfg,
			ReadTimeout:       0,
			ReadHeaderTimeout: 0,
			WriteTimeout:      0,
			IdleTimeout:       0,
			MaxHeaderBytes:    0,
			TLSNextProto:      nil,
			ConnState:         nil,
			ErrorLog:          nil,
			BaseContext:       nil,
			ConnContext:       nil,
		}

		wrappedGrpc := grpcweb.WrapServer(w.Handle())
		w.tlsHttpServer.Handler = http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			fmt.Printf("Handle request\n")
			if wrappedGrpc.IsGrpcWebRequest(req) {
				wrappedGrpc.ServeHTTP(resp, req)
				return
			}
			// Fall back to other servers.
			http.DefaultServeMux.ServeHTTP(resp, req)
		})
		lis, err := tls.Listen("tcp", w.BaseServer.Address(), &tlsCfg)
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
