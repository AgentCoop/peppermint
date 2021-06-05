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
	GetResponse() Response
	SendHeader()
}

func (p *calldesc) GetContext() context.Context {
	return p.Context
}

func (p *calldesc) GetRequest() Request {
	return p.Request
}

func (p *calldesc) GetResponse() Response {
	return p.Response
}

func (p *calldesc) SendHeader() {
	p.Context = metadata.NewOutgoingContext(p.Context, p.Request.MetaData())
}
