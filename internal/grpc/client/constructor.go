package client

import (
	"context"
	"google.golang.org/grpc"
	"net"
)

func NewBaseClient(addr net.Addr, opts ...grpc.DialOption) *baseClient {
	c := new(baseClient)
	c.addr = addr
	c.opts = opts
	return c
}

func NewBaseClientWithContext(ctx context.Context, addr net.Addr, opts ...grpc.DialOption) *baseClient {
	c := NewBaseClient(addr, opts...)
	c.ctx = ctx
	return c
}
