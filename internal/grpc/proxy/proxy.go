package proxy

import (
	"github.com/AgentCoop/peppermint/internal"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"google.golang.org/grpc"
)

type proxyStream struct {
	stream       interface{}
	nodeId       internal.NodeId
	encKey       []byte
	sentx, recvx int
}

type proxyConn struct {
	methodName string
	upstream   *proxyStream
	downstream *proxyStream
	downChan   chan g.CodecPacket
	upChan     chan g.CodecPacket
	upCallOpts []grpc.CallOption
}

func (p *proxyConn) WithUpstreamCallOptions(opts ...grpc.CallOption) {
	p.upCallOpts = opts
}
