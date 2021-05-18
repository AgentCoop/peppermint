package client

import (
	"context"
	"google.golang.org/grpc/metadata"
)

type reqResPair struct {
	context.Context
	Request
	Response
}

type RequestResponsePair interface {
	context.Context
	Request
	Response
	GetContext() context.Context
	GetRequest() Request
	AssignNewRequest(Request) Request
	GetResponse() Response
	AssignNewResponse(Response) Response
	SendHeader()
}

func NewRequestResponsePair(c BaseClient, ctx context.Context) *reqResPair {
	return &reqResPair{ctx,NewRequest(c), NewResponse(c, ctx)}
}

func (p *reqResPair) GetContext() context.Context {
	return p.Context
}

func (p *reqResPair) GetRequest() Request {
	return p.Request
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

func (p *reqResPair) SendHeader() {
	p.Context = metadata.NewOutgoingContext(p.Context, p.Request.MetaData())
}
