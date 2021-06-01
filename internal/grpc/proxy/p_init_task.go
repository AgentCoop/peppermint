package proxy

import (
	"context"
	job "github.com/AgentCoop/go-work"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (p *proxyConn) initTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		fullMethod := p.downstream.(runtime.StreamInfo).FullMethod()
		upstreamConn := j.GetValue().(*grpc.ClientConn)
		upstreamCtx, _ := context.WithCancel(p.downstream.Context())
		md, _ := metadata.FromIncomingContext(p.downstream.Context())
		upstreamCtx = metadata.NewOutgoingContext(upstreamCtx, md.Copy())

		upstream, err := grpc.NewClientStream(upstreamCtx, nil, upstreamConn, fullMethod)
		task.Assert(err)

		p.upstream = g.NewStream(upstream, fullMethod, p.upClient.EncKey())
	}
	run := func(task job.Task) {
		task.Done()
	}
	return init, run, nil
}

