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
			encKey := c.downstream.CallDesc().EncKey()
			newPacket := codec.NewRawPacket(p.Payload().([]byte), encKey)
			// Send response header from the backend to the client
			if c.downstream.ReceivedCount() == 0 {
				upCallDesc := c.upstream.CallDesc()
				upHeader := upCallDesc.Header()
				c.downstream.CallDesc().SetHeader(*upHeader)
			}
			c.downstream.Send(newPacket)
		case p := <-c.downstreamChan:
			task.Assert(p)
			encKey := c.upstream.CallDesc().EncKey()
			newPacket := codec.NewRawPacket(p.Payload().([]byte), encKey)
			c.upstream.Send(newPacket)
		}
		task.Tick()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

