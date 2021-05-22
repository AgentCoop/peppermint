package client

import (
	"context"
	job "github.com/AgentCoop/go-work"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
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
	SessionId() g.SessionId
	SetSessionId(id g.SessionId)
}

type baseClient struct {
	ctx context.Context
	conn    grpc.ClientConnInterface
	opts []grpc.DialOption
	address net.Addr
	onConnectedHook onConnectedHook
	withMiddlewareHook middlewareHook
	sId g.SessionId
}

func NewBaseClient(address net.Addr) *baseClient {
	c := new(baseClient)
	c.address = address
	return c
}

func NewBaseClientWithContext(address net.Addr, ctx context.Context, opts... grpc.DialOption) *baseClient {
	c := new(baseClient)
	c.address = address
	c.opts = opts
	return c
}

func (c *baseClient) SessionId() g.SessionId {
	return c.sId
}

func (c *baseClient) SetSessionId(id g.SessionId) {
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