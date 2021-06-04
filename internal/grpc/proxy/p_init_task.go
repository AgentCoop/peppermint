package proxy

import (
	job "github.com/AgentCoop/go-work"
	s "github.com/AgentCoop/peppermint/internal/grpc/stream"
	"google.golang.org/grpc"
)

func (p *proxyConn) initTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		fullMethod := p.downstream.FullMethod()
		upstreamConn := j.GetValue().(*grpc.ClientConn)
		upstream, err := grpc.NewClientStream(p.downstream.Context(), nil, upstreamConn, fullMethod)
		task.Assert(err)
		p.upstream = s.NewClientStream(upstream, p.upClient.IsSecure(), p.upClient.EncKey())
	}
	run := func(task job.Task) {
		task.Done()
	}
	return init, run, nil
}

