package balancer

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"google.golang.org/grpc"
)

type proxyConn struct {
	downstream grpc.ServerStream
	upstream grpc.ClientStream
	uprecvx int // upstream messages received count
	downrecvx int // downstream
	downEncKey []byte
	upEncKey []byte
	downstreamChan chan codec.Packet
	upstreamChan chan codec.Packet
}

func NewProxyConn(downstream grpc.ServerStream, upstream grpc.ClientStream, downKey []byte, upKey []byte) *proxyConn {
	conn := &proxyConn{
		downstream,
		upstream,
		0, 0,
		downKey, upKey,
		nil, nil,
	}
	return conn
}

func NewProxyConnJob() job.Job {
	j := job.NewJob(nil)
	proxyConn := NewProxyConn()
	j.AddTask(proxyConn.readDownstreamTask)
	j.AddTask(proxyConn.readUpstreamTask)
	j.AddTask(proxyConn.writeStreamTask)
	return j
}
