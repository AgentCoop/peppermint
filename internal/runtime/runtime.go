package runtime

import (
	i "github.com/AgentCoop/peppermint/internal"
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
	nodeCfg     NodeConfigurator
	parser      CliParser
	svcRegistry map[string]Service
}

func NewRuntime(nodeMngr NodeManager, nodeCfg NodeConfigurator, parser CliParser) *runtime {
	r := &runtime{
		nodeMngr: nodeMngr,
		nodeCfg:  nodeCfg,
		parser:   parser,
	}
	r.svcRegistry = make(map[string]Service, 0)
	return r
}

type Runtime interface {
	NodeManager() NodeManager
	CliParser() CliParser
	NodeConfigurator() NodeConfigurator
	RegisterService(string, Service)
	Services() []Service
	ServicePolicyByName(string) ServicePolicy
	ServiceByName(string) Service
}

func (r *runtime) RegisterService(svcName string, svc Service) {
	r.svcRegistry[svcName] = svc
}

func (r *runtime) ServiceByName(svcName string) Service {
	return r.svcRegistry[svcName]
}

func (r *runtime) ServicePolicyByName(svcName string) ServicePolicy {
	return r.ServiceByName(svcName).Policy()
}

func (r *runtime) Services() []Service {
	l := len(r.svcRegistry)
	out := make([]Service, l)
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

func (r *runtime) NodeConfigurator() NodeConfigurator {
	return r.nodeCfg
}

func (r *runtime) CliParser() CliParser {
	return r.parser
}
