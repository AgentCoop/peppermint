package proxy

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"google.golang.org/grpc"
)

func (c *proxyConn) writeStreamTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		select {
		case unpacker := <-c.upChan:
			nodeId := unpacker.NodeId()
			ss := c.downstream.stream.(grpc.ServerStream)
			task.Assert(p)
			encKey := c.downstream.CallDesc().EncKey()
			packer := codec.NewPacker(nodeId, unpacker.Payload(), unpacker.PayloadType(), c.downstream.encKey)
			// Send response header from the backend to the client
			if c.downstream.recvx == 0 {
				//upCallDesc := c.upstream.CallDesc()
				//upHeader := upCallDesc.Header()
				//c.downstream.CallDesc().SetHeader(*upHeader)
				//ss.S
			}
			ss.SendMsg(packer)
			//c.downstream.Send(newPacket)
		case p := <-c.downChan:
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

