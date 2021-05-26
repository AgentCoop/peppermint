package server

import (
	job "github.com/AgentCoop/go-work"
	_ "github.com/AgentCoop/peppermint/internal/grpc/codec"
	"google.golang.org/grpc"
	"net"
)

type BaseServer interface {
	Name() string
	Address() net.Addr
	Handle() *grpc.Server
	StartTask(j job.Job) (job.Init, job.Run, job.Finalize)
	RegisterServer()
}

type baseServer struct {
	name string
	address net.Addr
	task job.Task
	handle *grpc.Server
	lis net.Listener
}

func (s *baseServer) Address() net.Addr {
	return s.address
}

func (s *baseServer) RegisterServer() {
	panic("implement me")
}

func NewBaseServer(name string, address net.Addr, server *grpc.Server) *baseServer {
	s := new(baseServer)
	s.name = name
	s.address = address
	s.handle = server
	return s
}

func (s *baseServer) Handle() *grpc.Server {
	return s.handle
}

func (s *baseServer) Name() string {
	return s.name
}

func (s *baseServer) StartTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		s.task = task

		lis, err := net.Listen("tcp", s.address.String())
		task.Assert(err)

		s.lis = lis
	}
	run := func (task job.Task) {
		s.handle.Serve(s.lis)
		task.Done()
	}
	return init, run, nil
}
