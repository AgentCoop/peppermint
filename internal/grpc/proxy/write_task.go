package proxy

import (
	job "github.com/AgentCoop/go-work"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"google.golang.org/grpc"
)

func (c *proxyConn) writeStreamTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		ss := c.downstream.stream.(grpc.ServerStream)
		cs := c.upstream.stream.(grpc.ClientStream)
		var packet g.CodecPacket
		select {
		case upPacket := <-c.upChan:
			task.Assert(upPacket)
			if upPacket.HasFlag(g.PassthroughFlag) {
				packet = upPacket
			} else {
				payload, err := upPacket.Payload()
				task.Assert(err)
				packet = codec.NewPacket(c.downstream.nodeId, payload, c.downstream.encKey)
			}
			// Send response header from the backend to the client
			if c.downstream.sentx == 0 {
				header, err := cs.Header()
				task.Assert(err)
				ss.SendHeader(header)
			}
			ss.SendMsg(packet)
			c.downstream.sentx++
		case downPacket := <-c.downChan:
			task.Assert(downPacket)
			if downPacket.HasFlag(g.PassthroughFlag) {
				packet = downPacket
			} else {
				payload, err := downPacket.Payload()
				task.Assert(err)
				packet = codec.NewPacket(c.upstream.nodeId, payload, c.upstream.encKey)
			}
			cs.SendMsg(packet)
			c.upstream.sentx++
		}
		task.Tick()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

