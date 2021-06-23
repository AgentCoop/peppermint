package proxy

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"io"
)

func (c *proxyConn) readUpstreamTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		var err error
		desc := c.upstream.CallDesc()
		nodeId := desc.Meta().NodeId()
		packet := codec.NewPacket(nodeId, codec.RawPacket, desc.SecPolicy().EncKey())
		err = c.upstream.RecvMsg(packet)
		task.Assert(err)
		if err == io.EOF {
			task.Done()
			return
		}
		c.upChan <- packet
		task.Tick()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}


