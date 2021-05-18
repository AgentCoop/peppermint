package client

import (
	job "github.com/AgentCoop/go-work"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"google.golang.org/grpc"
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
	conn    grpc.ClientConnInterface
	address string
	onConnectedHook onConnectedHook
	withMiddlewareHook middlewareHook
	sId g.SessionId
}

func NewBaseClient(address string) *baseClient {
	c := new(baseClient)
	c.address = address
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
		var opts []grpc.DialOption
		opts = append(opts,
			grpc.WithInsecure(),
			c.withMiddlewareHook(),
			grpc.WithDefaultCallOptions(grpc.CallContentSubtype(codec.Name)))
		conn, err := grpc.Dial(c.address, opts...)
		task.Assert(err)
		c.conn = conn
		c.onConnectedHook(conn)
		task.Done()
	}
	return nil, run, nil
}
