package balancer

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
			newPacket := codec.NewRawPacket(p.Payload().([]byte), c.downEncKey)
			// Send response header from the backend to the client
			if c.uprecvx == 0 {
				md, _ := c.upstream.Header()
				c.downstream.SetHeader(md)
			}
			c.downstream.SendMsg(newPacket)
		case p := <-c.downstreamChan:
			task.Assert(p)
			newPacket := codec.NewRawPacket(p.Payload().([]byte), c.upEncKey)
			c.upstream.SendMsg(newPacket)
		}
		task.Tick()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

