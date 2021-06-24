package balancer

import (
	"github.com/AgentCoop/peppermint/internal/utils"
	"github.com/AgentCoop/peppermint/pkg/service/balancer"
	"google.golang.org/grpc"
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
	switch algo {
	case balancer.Random:
	case balancer.RoundRobin:
	}
	return nil
}
