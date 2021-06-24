package balancer

import (
	"google.golang.org/grpc"
)

type Balancer interface {
	ForwardCall(srv interface{}, stream grpc.ServerStream) error
	//SimpleRandom(svcName string, pool runtime.NodePool) runtime.ServiceEndpoint
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
