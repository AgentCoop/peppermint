package proxy

import (
	job "github.com/AgentCoop/go-work"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"google.golang.org/grpc"
	"io"
)

func (c *proxyConn) readDownstreamTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		var err error
		ss := c.upstream.stream.(grpc.ServerStream)
		packet := codec.NewPacket(0, nil, c.downstream.encKey)
		if len(c.upstream.encKey) > 0 && len(c.downstream.encKey) > 0 {
			packet.WithFlags(g.PassthroughFlag)
		}
		err = ss.RecvMsg(packet)
		task.Assert(err)
		if err == io.EOF {
			task.Done()
			return
		}
		c.downChan <- packet
		task.Tick()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

