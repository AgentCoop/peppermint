package runtime

import (
	"github.com/AgentCoop/peppermint/internal/runtime/store"
	"github.com/AgentCoop/peppermint/pkg"
	"github.com/AgentCoop/peppermint/pkg/service"
)

func NewRuntime(nodeMngr pkg.NodeManager, parser pkg.CliParser) *runtime {
	r := &runtime{
		nodeMngr: nodeMngr,
		parser:   parser,
	}
	r.svcRegistry = make(map[string]service.Service, 0)
	r.encKeyStore = store.NewInMemoryStore()
	return r
}
