package client

import (
	"context"
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc"
	"net"
)

type ReqChan chan ClientCallDescriptor
type ResChan chan struct{}
type onConnectedHook func(grpc.ClientConnInterface)

type BaseClient interface {
	ConnectTask(j job.Job) (job.Init, job.Run, job.Finalize)
	Connection() grpc.ClientConnInterface
	OnConnectedHook(onConnectedHook)
	WithUnaryInterceptors(...grpc.UnaryClientInterceptor)
	NodeId() i.NodeId
	IsSecure() bool
	EncKey() []byte
	SessionId() i.SessionId
	SetSessionId(id i.SessionId)
}

type baseClient struct {
	ctx               context.Context
	unaryInterceptors []grpc.UnaryClientInterceptor
	encKey            []byte
	conn              grpc.ClientConnInterface
	opts              []grpc.DialOption
	address           net.Addr
	onConnectedHook   onConnectedHook
	sId               i.SessionId
}

func NewBaseClient(endpoint runtime.ServiceEndpoint, opts ...grpc.DialOption) *baseClient {
	c := new(baseClient)
	c.address = endpoint.Address()
	c.encKey = endpoint.EncKey()
	c.opts = opts
	return c
}

func NewBaseClientWithContext(ctx context.Context, endpoint runtime.ServiceEndpoint, opts ...grpc.DialOption) *baseClient {
	c := NewBaseClient(endpoint, opts...)
	c.ctx = ctx
	return c
}

func (c *baseClient) IsSecure() bool {
	return false
}

func (c *baseClient) WithUnaryInterceptors(interceptors ...grpc.UnaryClientInterceptor) {
	c.unaryInterceptors = interceptors
}

func (c *baseClient) SessionId() i.SessionId {
	return c.sId
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

func (c *baseClient) OnConnectedHook(hook onConnectedHook) {
	c.onConnectedHook = hook
}
