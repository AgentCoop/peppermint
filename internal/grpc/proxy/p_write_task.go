package proxy

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc"
)

func (c *proxyConn) writeStreamTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		select {
		case p := <-c.upstreamChan:
			task.Assert(p)
			downStreamInfo := c.downstream.(runtime.StreamInfo)
			upStreamInfo := c.upstream.(runtime.StreamInfo)
			newPacket := codec.NewRawPacket(p.Payload().([]byte), downStreamInfo.EncKey())
			// Send response header from the backend to the client
			if upStreamInfo.MessagesReceived() == 0 {
				md, _ := c.upstream.(grpc.ClientStream).Header()
				c.downstream.(grpc.ServerStream).SetHeader(md)
			}
			c.downstream.SendMsg(newPacket)
		case p := <-c.downstreamChan:
			task.Assert(p)
			upStreamInfo := c.upstream.(runtime.StreamInfo)
			newPacket := codec.NewRawPacket(p.Payload().([]byte),upStreamInfo.EncKey())
			c.upstream.SendMsg(newPacket)
		}
		task.Tick()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

