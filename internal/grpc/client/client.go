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

type Client interface {
	GetServerAddress() string
	GetNodeId() []byte
	//GetDefaultHeader() *service.RequestHeader
	Connect()
}

type BaseClient struct {
	Conn *grpc.ClientConn
	Task job.Task
	Address string
}

func (c *BaseClient) Connect() {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(c.UnaryClientInterceptor()),
		grpc.WithDefaultCallOptions(grpc.CallContentSubtype(codec.Name)))
	conn, err := grpc.Dial(c.Address, opts...)
	c.Task.Assert(err)
	c.Conn = conn
}

func (c *BaseClient) ConnectTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		var opts []grpc.DialOption
		opts = append(opts,
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(c.UnaryClientInterceptor()),
			grpc.WithDefaultCallOptions(grpc.CallContentSubtype(codec.Name)))
		conn, err := grpc.Dial(c.Address, opts...)
		task.Assert(err)
		c.Conn = conn
		task.Done()
	}
	return nil, run, nil
}

func (c *BaseClient) UnaryClientInterceptor(opts ...grpc.CallOption) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx = c.addMetaData(ctx)
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}

func (c *BaseClient) addMetaData(ctx context.Context) context.Context {
	send, _ := metadata.FromOutgoingContext(ctx)
	newMD := metadata.Pairs(META_NODE_ID, "v3")
	return metadata.NewOutgoingContext(ctx, metadata.Join(send, newMD))
}
