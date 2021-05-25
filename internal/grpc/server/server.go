package server

import (
	"context"
	job "github.com/AgentCoop/go-work"
	_ "github.com/AgentCoop/peppermint/internal/grpc/codec"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc"
	"net"
)

type PairChan chan RequestResponsePair

type MetaData interface {
	context.Context
	RequestHeader
	ResponseHeader
}

type reqResPair struct {
	context.Context
	Request
	Response
	runtime.Configurator
}

type RequestResponsePair interface {
	context.Context
	Request
	Response
	GetRequest() Request
	AssignNewRequest(Request) Request
	GetResponse() Response
	AssignNewResponse(Response) Response
	GetConfigurator() runtime.Configurator
}

func NewRequestResponsePair(ctx context.Context, cfg runtime.Configurator) *reqResPair {
	return &reqResPair{ctx,NewRequest(ctx), NewResponse(ctx), cfg}
}

func (p *reqResPair) GetRequest() Request {
	return p.Request
}

func (p *reqResPair) GetConfigurator() runtime.Configurator {
	return p.Configurator
}

// Replaces the base request with an extended one
func (p *reqResPair) AssignNewRequest(new Request) Request {
	base := p.Request
	p.Request = new
	return base
}

func (p *reqResPair) GetResponse() Response {
	return p.Response
}

// Replaces the base response with an extended one
func (p *reqResPair) AssignNewResponse(new Response) Response {
	base := p.Response
	p.Response = new
	return base
}

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
