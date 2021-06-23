package proxy

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"google.golang.org/grpc"
	"io"
)

func (c *proxyConn) readUpstreamTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		var err error
		var packet g.CodecPacket
		cs := c.upstream.stream.(grpc.ClientStream)
		packet = codec.NewPacket(0, nil, c.upstream.encKey)
		// Do nothing with a packet. Will be unpacked by the downstream node
		if len(c.upstream.encKey) > 0 && len(c.downstream.encKey) > 0 {
			packet.WithFlags(g.PassthroughFlag)
		}
		err = cs.RecvMsg(packet)
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


