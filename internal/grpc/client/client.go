package client

import (
	"context"
	job "github.com/AgentCoop/go-work"

	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	META_NODE_ID = "node_id"
)

type ReqChan chan Request
type ResChan chan Response
type onConnectedHook func(grpc.ClientConnInterface)


type BaseClient interface {
	ConnectTask(j job.Job) (job.Init, job.Run, job.Finalize)
	Connection() grpc.ClientConnInterface
	OnConnectedHook(hook onConnectedHook)
}

type baseClient struct {
	conn    grpc.ClientConnInterface
	address string
	onConnectedHook onConnectedHook
}

func NewBaseClient(address string) *baseClient {
	c := new(baseClient)
	c.address = address
	return c
}

func (c *baseClient) Connection() grpc.ClientConnInterface {
	return c.conn
}

func (c *baseClient) OnConnectedHook(hook onConnectedHook)  {
	c.onConnectedHook = hook
}

func (c *baseClient) ConnectTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		var opts []grpc.DialOption
		opts = append(opts,
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(c.UnaryClientInterceptor()),
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
