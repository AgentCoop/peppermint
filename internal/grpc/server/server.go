package server

import (
	job "github.com/AgentCoop/go-work"
	_ "github.com/AgentCoop/peppermint/internal/grpc/codec"
	"net"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type GrpcServer interface {
	GetTask() job.Task
	StartTask(j job.Job) (job.Init, job.Run, job.Finalize)
	RegisterService()
	Session() Session
}

type BaseServer struct {
	Address string
	task job.Task
	Handle *grpc.Server
	lis net.Listener
	session Session
}

func DefaultGrpcServer() *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(UnaryServerInterceptor()),
	}
	return grpc.NewServer(opts...)
}

func (s *BaseServer) Session() Session {
	return s.session
}

func (s *BaseServer) GetTask() job.Task {
	return s.task
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		p, _ := peer.FromContext(ctx)
		wrappedReq := &Request{}
		wrappedReq.original = req
		wrappedReq.clientAddr = p.Addr
		return handler(ctx, wrappedReq)
	}
}

func (s *BaseServer) StartTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		s.task = task

		lis, err := net.Listen("tcp", s.Address)
		task.Assert(err)

		s.lis = lis
		//hub.RegisterHubServer(s.grpc, s)
	}
	run := func (task job.Task) {
		s.Handle.Serve(s.lis)
		task.Done()
	}
	return init, run, nil
}
