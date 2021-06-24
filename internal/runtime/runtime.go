package runtime

import (
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/pkg"
	"github.com/AgentCoop/peppermint/pkg/node"
	rt "github.com/AgentCoop/peppermint/pkg/runtime"
	"github.com/AgentCoop/peppermint/pkg/service"
	"net"
)

type NodeStatus int

const (
	Available NodeStatus = iota + 1
)

type NodePool interface {
	Add(node.Node)
	Remove(i.NodeId)
	FindById(i.NodeId) node.Node
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
	nodeMngr    node.NodeManager
	nodeCfg     node.NodeConfigurator
	parser      pkg.CliParser
	svcRegistry map[string]service.Service
	encKeyStore rt.InMemoryStore
}

func (r *runtime) RegisterService(svcName string, svc service.Service) {
	r.svcRegistry[svcName] = svc
}

func (r *runtime) ServiceByName(svcName string) service.Service {
	return r.svcRegistry[svcName]
}

func (r *runtime) ServicePolicyByName(svcName string) service.ServicePolicy {
	svc := r.ServiceByName(svcName)
	if svc == nil {
		return nil
	}
	return svc.Policy()
}

func (r *runtime) Services() []service.Service {
	l := len(r.svcRegistry)
	out := make([]service.Service, l)
	i := 0
	for _, svc := range r.svcRegistry {
		out[i] = svc
		i++
	}
	return out
}

func (r *runtime) NodeManager() node.NodeManager {
	return r.nodeMngr
}

func (r *runtime) NodeConfigurator() node.NodeConfigurator {
	return r.nodeCfg
}

func (r *runtime) CliParser() pkg.CliParser {
	return r.parser
}

func (m *runtime) EncKeyStore() rt.InMemoryStore {
	return m.encKeyStore
}