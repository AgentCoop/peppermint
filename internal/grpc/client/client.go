package client

import (
	"context"
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	"google.golang.org/grpc"
	"net"
)

type connProvider func(grpc.ClientConnInterface)

type BaseClient interface {
	Context() context.Context
	ConnectTask(j job.Job) (job.Init, job.Run, job.Finalize)
	Connection() grpc.ClientConnInterface
	WithContext(ctx context.Context)
	WithConnProvider(connProvider)
	WithEncKey([]byte)
	WithUnaryInterceptors(...grpc.UnaryClientInterceptor)
	NodeId() i.NodeId
	IsSecure() bool
	EncKey() []byte
	SessionId() i.SessionId
	SetSessionId(id i.SessionId)
}

type baseClient struct {
	ctx               context.Context
	addr              net.Addr
	opts              []grpc.DialOption
	conn              grpc.ClientConnInterface
	connProvider      connProvider
	unaryInterceptors []grpc.UnaryClientInterceptor
	encKey            []byte
	sId               i.SessionId
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

func (c *baseClient) WithConnProvider(provider connProvider) {
	c.connProvider = provider
}

func (c *baseClient) SessionId() i.SessionId {
	return c.sId
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
