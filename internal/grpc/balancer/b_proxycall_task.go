package balancer

import (
	"context"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (b *balancerCtx) proxyCallTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		ctx.Request = server.NewRequest(stream.Context())
		ctx.upstreamCtx, ctx.upstreamCancel = context.WithCancel(stream.Context())
		md, _ := metadata.FromIncomingContext(stream.Context())

		ctx.fullMethod, _ = grpc.MethodFromServerStream(b.stream)
		b.Request = server.NewRequest(b.stream.Context())
		b.upstreamCtx, b.upstreamCancel = context.WithCancel(b.stream.Context())
		md, _ := metadata.FromIncomingContext(b.stream.Context())
		b.upstreamCtx = metadata.NewOutgoingContext(b.upstreamCtx, md.Copy())
	}
	run := func(task job.Task) {
		fullMethodName, methodErr := grpc.MethodFromServerStream(b.stream)
		task.Assert(methodErr)
		upstreamConn := j.GetValue().(*grpc.ClientConn)
		b.setupUpstreamHeader()
		upstream, err := grpc.NewClientStream(b.upstreamCtx, nil, upstreamConn, fullMethodName)
		task.Assert(err)
		pconn := NewProxyConn(b.stream, upstream, nil, nil)
		// Create and run a job to forward current gRPC call
		proxyJob := job.NewJob(nil)
		proxyJob.AddTask(pconn.readUpstreamTask)
		proxyJob.AddTask(pconn.readDownstreamTask)
		proxyJob.AddTask(pconn.writeStreamTask)
		<-proxyJob.Run()
		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

