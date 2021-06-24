package service

import (
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc"
)

type HubService interface {
	runtime.Service
}

type WebProxyBalancer interface {
	ForwardCall(srv interface{}, stream grpc.ServerStream) error
	SimpleRandom(svcName string, pool runtime.NodePool) runtime.ServiceEndpoint
	//RoundRobin(string, runtime.NodePool) runtime.ServiceEndpoint
	///LeastConns(string, runtime.NodePool) runtime.ServiceEndpoint
}
