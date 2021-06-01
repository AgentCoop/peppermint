package client

import (
	"context"
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc"
	"net"
)

type ReqChan chan RequestResponsePair
type ResChan chan struct{}

type onConnectedHook func(grpc.ClientConnInterface)
type middlewareHook func() grpc.DialOption

type BaseClient interface {
	ConnectTask(j job.Job) (job.Init, job.Run, job.Finalize)
	Connection() grpc.ClientConnInterface
	OnConnectedHook(onConnectedHook)
	WithMiddlewareHook(middlewareHook)
	SessionId() i.SessionId
	SetSessionId(id i.SessionId)
	EncKey() []byte
}

type baseClient struct {
	ctx context.Context
	encKey []byte
	conn    grpc.ClientConnInterface
	opts []grpc.DialOption
	address net.Addr
	onConnectedHook onConnectedHook
	withMiddlewareHook middlewareHook
	sId i.SessionId
}

func NewBaseClient(endpoint runtime.ServiceEndpoint, opts... grpc.DialOption) *baseClient {
	c := new(baseClient)
	c.address = endpoint.Address()
	c.encKey = endpoint.EncKey()
	c.opts = opts
	return c
}

func NewBaseClientWithContext(ctx context.Context, endpoint runtime.ServiceEndpoint, opts... grpc.DialOption) *baseClient {
	c := NewBaseClient(endpoint, opts...)
	c.ctx = ctx
	return c
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

func (c *baseClient) OnConnectedHook(hook onConnectedHook)  {
	c.onConnectedHook = hook
}

func (c *baseClient) WithMiddlewareHook(hook middlewareHook) {
	c.withMiddlewareHook = hook
}

func (c *baseClient) ConnectTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		var opts = []grpc.DialOption{
			grpc.WithInsecure(),
			c.withMiddlewareHook(),
			grpc.WithDefaultCallOptions(grpc.CallContentSubtype(codec.Name)),
		}
		var conn *grpc.ClientConn
		var err error
		switch {
		case c.ctx != nil:
			conn, err = grpc.DialContext(c.ctx, c.address.String(), opts...)
		default:
			conn, err = grpc.Dial(c.address.String(), opts...)
		}
		task.Assert(err)
		c.conn = conn
		c.onConnectedHook(conn)
		task.Done()
	}
	return nil, run, nil
}
