package proxy

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"io"
)

func (c *proxyConn) readDownstreamTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		var err error
		downStreamInfo := c.downstream.(runtime.StreamInfo)
		recvRaw := codec.NewRawPacket(nil, downStreamInfo.EncKey())
		err = c.downstream.RecvMsg(recvRaw)
		task.Assert(err)
		if err == io.EOF {
			task.Done()
			return
		}
		c.downstreamChan <- recvRaw
		//c.downrecvx++
		task.Tick()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

