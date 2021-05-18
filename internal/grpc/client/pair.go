package client

import "context"

type reqResPair struct {
	context.Context
	Request
	Response
}

type RequestResponsePair interface {
	context.Context
	Request
	Response
	GetRequest() Request
	AssignNewRequest(Request) Request
	GetResponse() Response
	AssignNewResponse(Response) Response
}

func NewRequestResponsePair(c BaseClient, ctx context.Context) *reqResPair {
	return &reqResPair{ctx,NewRequest(c, ctx), NewResponse(c, ctx)}
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
