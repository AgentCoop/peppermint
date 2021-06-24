package balancer

import (
	"github.com/AgentCoop/peppermint/internal/utils"
	"github.com/AgentCoop/peppermint/pkg/service/balancer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

func (lb *lbService) ForwardCall(srv interface{}, stream grpc.ServerStream) error {
	methodName, ok := grpc.MethodFromServerStream(stream)
	svcName := utils.Grpc_MethodToServiceName(methodName)
	if !ok {
		return nil
	}
	cfg := lb.Configurator().(balancer.BalancerConfigurator)
	algo, ok := cfg.PreferredAlgoByServiceName(svcName)
	if !ok {
		algo = cfg.DefaultAlgo()
	}
	var addr net.Addr
	switch algo {
	case balancer.Random:
		addr = lb.RandomChoice(svcName)
	case balancer.RoundRobin:
		addr = lb.RoundRobinChoice(svcName)
	}
	if addr == nil {
		return status.Errorf(codes.Unavailable, "requested service %s temporarily unavailable", svcName)
	}
	return nil
}
