package client

import (
	"context"
	job "github.com/AgentCoop/go-work"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/utils"

	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	//middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/client"
	//md_middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/client/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	META_NODE_ID = "node_id"
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


//	GetContext()
	//CreateRequest() Request
	//HandleResponse(context.Context)
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

func (c *baseClient) ParseResponseHeader(ctx context.Context) ResponseHeader {
	md, ok := metadata.FromIncomingContext(ctx)
	md2, _ := metadata.FromOutgoingContext(ctx)
	_ = md2
	if ! ok {
		//panic("metadata")
	}
	sId := utils.GetSessionId(&md)
	c.SetSessionId(sId)
	return nil
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

func (c *baseClient) UnaryClientInterceptor(opts ...grpc.CallOption) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx = c.addMetaData(ctx)
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}

func (c *baseClient) addMetaData(ctx context.Context) context.Context {
	send, _ := metadata.FromOutgoingContext(ctx)
	newMD := metadata.Pairs(META_NODE_ID, "v3")
	return metadata.NewOutgoingContext(ctx, metadata.Join(send, newMD))
}
