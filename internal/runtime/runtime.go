package runtime

import (
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/runtime/deps"
	//"github.com/AgentCoop/peppermint/internal/service"
	"net"
)

type NodeStatus int

const (
	Available NodeStatus = iota + 1
)

type NodePool interface {
	Add(Node)
	Remove(i.NodeId)
	FindById(i.NodeId) Node
	FilterByStatus(NodeStatus) NodePool
	Len() int
}

type ServiceLocator interface {
	FindByMethodName(string) NodePool
	ServiceNameByMethod(string) string
}

type ServiceEndpoint interface {
	Address() net.Addr
	EncKey() []byte
}

type runtime struct {
	nodeMngr    NodeManager
	nodeCfg     deps.NodeConfigurator
	parser      deps.CliParser
	svcRegistry map[string]grpc.Service
}

func NewRuntime(
	nodeMngr NodeManager,
	nodeCfg deps.NodeConfigurator,
	parser deps.CliParser,
) *runtime {
	r := &runtime{
		nodeMngr: nodeMngr,
		nodeCfg:  nodeCfg,
		parser:   parser,
	}
	return r
}

type Runtime interface {
	NodeManager() NodeManager
	CliParser() deps.CliParser
	NodeConfigurator() deps.NodeConfigurator
	RegisterService(string, grpc.Service)
	Services() []grpc.Service
	ServicePolicyByName(string) grpc.ServicePolicy
}

func (r *runtime) RegisterService(svcName string, svc grpc.Service) {
	r.svcRegistry[svcName] = svc
}

func (r *runtime) ServicePolicyByName(svcName string) grpc.ServicePolicy {
	return r.svcRegistry[svcName].Policy()
}

func (r *runtime) Services() []grpc.Service {
	l := len(r.svcRegistry)
	out := make([]grpc.Service, l)
	i := 0
	for _, svc := range r.svcRegistry {
		out[i] = svc
		i++
	}
	return out
}

func (r *runtime) NodeManager() NodeManager {
	return r.nodeMngr
}

func (r *runtime) NodeConfigurator() deps.NodeConfigurator {
	return r.nodeCfg
}

func (r *runtime) CliParser() deps.CliParser {
	return r.parser
}
