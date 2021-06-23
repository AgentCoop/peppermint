package proxy

import (
	"context"
	job "github.com/AgentCoop/go-work"
	"google.golang.org/grpc"
)

func (p *proxyConn) initTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		upstreamConn := j.GetValue().(*grpc.ClientConn)
		streamDesc := &grpc.StreamDesc{
			StreamName:    "",
			Handler:       nil,
			ServerStreams: false,
			ClientStreams: false,
		}
		cs, err := grpc.NewClientStream(
			context.Background(),
			streamDesc,
			upstreamConn,
			p.methodName,
			p.upCallOpts...,
		)
		task.Assert(err)
		p.upstream = &proxyStream{
			stream: cs,
			nodeId: 0,
			encKey: nil,
		}
	}
	run := func(task job.Task) {
		task.Done()
	}
	return init, run, nil
}

