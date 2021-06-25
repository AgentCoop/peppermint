package client

import (
	"context"
	i "github.com/AgentCoop/peppermint/internal"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"google.golang.org/grpc"
)

type baseClient struct {
	ctx               context.Context
	target            string
	opts              []grpc.DialOption
	conn              grpc.ClientConnInterface
	connProvider      g.ConnProvider
	timeoutMs         uint
	port              uint16
	unaryInterceptors []grpc.UnaryClientInterceptor
	encKey            []byte
	sId               i.SessionId
	lastCall          g.ClientDescriptor
}

func (c *baseClient) Context() context.Context {
	switch {
	case c.ctx == nil:
		return context.Background()
	default:
		return c.ctx
	}
}

func (c *baseClient) NodeId() i.NodeId {
	panic("implement me")
}

func (c *baseClient) IsSecure() bool {
	return false
}

func (c *baseClient) WithContext(ctx context.Context) {
	c.ctx = ctx
}

func (c *baseClient) WithUnaryInterceptors(interceptors ...grpc.UnaryClientInterceptor) {
	c.unaryInterceptors = interceptors
}

func (c *baseClient) WithConnProvider(provider g.ConnProvider) {
	c.connProvider = provider
}

func (c *baseClient) SessionId() i.SessionId {
	return c.sId
}

func (c *baseClient) WithTimeout(ms uint) {
	c.timeoutMs = ms
}

func (c *baseClient) WithTargetPort(port uint16) {
	c.port = port
}

func (c *baseClient) WithEncKey(key []byte) {
	c.encKey = key
}

func (c *baseClient) EncKey() []byte {
	return c.encKey
}

func (c *baseClient) SetSessionId(id i.SessionId) {
	c.sId = id
}

func (c *baseClient) Connection() grpc.ClientConnInterface {
	return c.conn
}

func (c *baseClient) LastCall() g.ClientDescriptor {
	return c.lastCall
}

func (c *baseClient) SetLastCall(call g.ClientDescriptor) {
	c.lastCall = call
}
