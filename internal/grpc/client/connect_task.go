package client

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"google.golang.org/grpc"
)

func (c *baseClient) ConnectTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		var opts = []grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithChainUnaryInterceptor(c.unaryInterceptors...),
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
		c.connProvider(conn)
		task.Done()
	}
	return nil, run, nil
}

