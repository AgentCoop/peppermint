package server

import (
	"context"
	job "github.com/AgentCoop/go-work"
	_ "github.com/AgentCoop/peppermint/internal/grpc/codec"
	"google.golang.org/grpc"
	"net"
)

type MetaData interface {
	context.Context
	RequestHeader
	ResponseHeader
}

type BaseServer interface {
	Handle() *grpc.Server
	StartTask(j job.Job) (job.Init, job.Run, job.Finalize)
	RegisterServer()
}

type baseServer struct {
	address string
	task job.Task
	handle *grpc.Server
	lis net.Listener
}

func (s *baseServer) RegisterServer() {
	panic("implement me")
}

func NewBaseServer(address string, server *grpc.Server) *baseServer {
	s := new(baseServer)
	s.address = address
	s.handle = server
	return s
}

func (s *baseServer) Handle() *grpc.Server {
	return s.handle
}

func (s *baseServer) GetTask() job.Task {
	return s.task
}

func (s *baseServer) StartTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		s.task = task

		lis, err := net.Listen("tcp", s.address)
		task.Assert(err)

		s.lis = lis
	}
	run := func (task job.Task) {
		s.handle.Serve(s.lis)
		task.Done()
	}
	return init, run, nil
}
