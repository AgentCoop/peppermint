package server

import (
	"crypto/tls"
	job "github.com/AgentCoop/go-work"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"net/http"
)

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
		addr := w.BaseServer.Address().String()
		lis, err := tls.Listen("tcp", addr, &tlsCfg)
		task.Assert(err)
		w.lis = lis
		//j.Log(0, w.BaseServer.Name()) <- fmt.Sprintf("is listening on %s", addr)
	}
	run := func(task job.Task) {
		w.tlsHttpServer.Serve(w.lis)
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

