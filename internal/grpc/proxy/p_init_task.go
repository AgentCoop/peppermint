package proxy

import (
	job "github.com/AgentCoop/go-work"
	s "github.com/AgentCoop/peppermint/internal/grpc/stream"
	"google.golang.org/grpc"
)

func (p *proxyConn) initTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		desc := p.downstream.CallDesc()
		upstreamConn := j.GetValue().(*grpc.ClientConn)
		streamDesc := &grpc.StreamDesc{
			StreamName:    "",
			Handler:       nil,
			ServerStreams: false,
			ClientStreams: false,
		}
		cs, err := grpc.NewClientStream(
			p.downstream.Context(),
			streamDesc,
			upstreamConn,
			desc.Method().FullName(),
			p.upCallOpts...,
		)
		task.Assert(err)
		p.upstream = s.NewClientStream(cs, nil)
	}
	run := func(task job.Task) {
		task.Done()
	}
	return init, run, nil
}

