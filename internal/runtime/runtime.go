package runtime

import (
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/pkg"
	rt "github.com/AgentCoop/peppermint/pkg/runtime"
	"github.com/AgentCoop/peppermint/pkg/service"
	"net"
)

type NodeStatus int

const (
	Available NodeStatus = iota + 1
)

type NodePool interface {
	Add(pkg.Node)
	Remove(i.NodeId)
	FindById(i.NodeId) pkg.Node
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
	app         pkg.App
	nodeMngr    pkg.NodeManager
	parser      pkg.CliParser
	svcRegistry map[string]service.Service
	encKeyStore rt.InMemoryStore
}

func (r *runtime) App() pkg.App {
	return r.app
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

func (r *runtime) NodeManager() pkg.NodeManager {
	return r.nodeMngr
}

func (r *runtime) CliParser() pkg.CliParser {
	return r.parser
}

func (m *runtime) EncKeyStore() rt.InMemoryStore {
	return m.encKeyStore
}
