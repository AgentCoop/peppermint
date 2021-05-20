package balancer

import "google.golang.org/grpc"

type proxyLink struct {
	ss grpc.ServerStream
	cs grpc.ClientStream
}

type grpcMessage struct {

}

type proxyConn struct {
	upstream proxyLink
	downstream proxyLink
	msgReceived int
	msgSent int
}

func NewProxyLink() *proxyLink {
	return nil
}
