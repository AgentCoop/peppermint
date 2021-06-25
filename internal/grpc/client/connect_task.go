package client

import (
	"context"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"google.golang.org/grpc"
	"time"
)

func (c *baseClient) ConnectTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		var opts = []grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithChainUnaryInterceptor(c.unaryInterceptors...),
			grpc.WithDefaultCallOptions(grpc.CallContentSubtype(codec.Name)),
		}
		ctx := context.Background()
		if c.timeoutMs > 0 {
			deadline := time.Now().Add(time.Duration(c.timeoutMs) * time.Millisecond)
			ctx, _ = context.WithDeadline(ctx, deadline)
		}
		conn, err := grpc.DialContext(ctx, c.addr.String(), opts...)
		task.Assert(err)
		c.conn = conn
		c.connProvider(conn)
		task.Done()
	}
	return nil, run, nil
}
