package client

import (
	"context"
	"google.golang.org/grpc/metadata"
)

type calldesc struct {
	context.Context
	Request
	Response
}

type ClientCallDescriptor interface {
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

func (p *calldesc) GetContext() context.Context {
	return p.Context
}

func (p *calldesc) GetRequest() Request {
	return p.Request
}

// Replaces the base request with an extended one
func (p *calldesc) AssignNewRequest(new Request) Request {
	base := p.Request
	p.Request = new
	return base
}

func (p *calldesc) GetResponse() Response {
	return p.Response
}

// Replaces the base response with an extended one
func (p *calldesc) AssignNewResponse(new Response) Response {
	base := p.Response
	p.Response = new
	return base
}

func (p *calldesc) SendHeader() {
	p.Context = metadata.NewOutgoingContext(p.Context, p.Request.MetaData())
}
