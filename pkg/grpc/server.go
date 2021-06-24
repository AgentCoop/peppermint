package grpc

import (
	job "github.com/AgentCoop/go-work"
	"google.golang.org/grpc"
	"net"
)

type BaseServer interface {
	FullName() string
	Methods() []string
	Address() net.Addr
	Handle() *grpc.Server
	StartTask(j job.Job) (job.Init, job.Run, job.Finalize)
	RegisterServer()
	WithStdoutLogger(handler job.LogHandler)
}
