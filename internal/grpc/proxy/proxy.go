package proxy

import (
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"google.golang.org/grpc"
)

type proxyStream struct {
	stream       interface{}
	encKey       []byte
	sentx, recvx int
}

type proxyConn struct {
	stripDown  bool // Unpack usptream packeted data if any
	upstream   *proxyStream
	downstream *proxyStream
	downChan   chan g.CodecPacker
	upChan     chan g.CodecUnpacker
	upCallOpts []grpc.CallOption
}

func (p *proxyConn) WithUpstreamCallOptions(opts ...grpc.CallOption) {
	p.upCallOpts = opts
}
