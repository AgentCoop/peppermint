package node

import "github.com/AgentCoop/peppermint/internal/grpc"

type Node interface {
	Id() grpc.NodeId
	ServiceAddressByName(string) grpc.ServiceAddress
}
