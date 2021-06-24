package balancer

import (
	"google.golang.org/grpc"
	"net"
)

type Balancer interface {
	ForwardCall(srv interface{}, stream grpc.ServerStream) error
	RandomChoice(svcName string) net.Addr
	RoundRobinChoice(svcName string) net.Addr
}

type Algo int

const (
	Random Algo = iota
	RoundRobin
)

func (algo Algo) String() string {
	return [...]string{"random", "round-robin"}[algo]
}

type BalancerConfigurator interface {
	PreferredAlgoByServiceName(svcName string) (Algo, bool)
	DefaultAlgo() Algo
}
