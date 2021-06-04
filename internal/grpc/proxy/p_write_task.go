package proxy

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
)

func (c *proxyConn) writeStreamTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		select {
		case p := <-c.upstreamChan:
			task.Assert(p)
			newPacket := codec.NewRawPacket(p.Payload().([]byte), c.downstream.EncKey())
			// Send response header from the backend to the client
			if c.downstream.MessagesReceived() == 0 {
				upHeader := c.upstream.Header()
				c.downstream.WithNewHeader(&upHeader)
			}
			c.downstream.Send(newPacket)
		case p := <-c.downstreamChan:
			task.Assert(p)
			newPacket := codec.NewRawPacket(p.Payload().([]byte), c.upstream.EncKey())
			c.upstream.Send(newPacket)
		}
		task.Tick()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

