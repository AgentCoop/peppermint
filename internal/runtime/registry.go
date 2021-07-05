package runtime

import (
	"github.com/AgentCoop/peppermint/pkg"
	rt "github.com/AgentCoop/peppermint/pkg/runtime"
)

type regKey string
type registryMap map[regKey][]interface{}

type Hook int
type HookHandler func(...interface{}) error
type hookEntry struct {
	typ     Hook
	handler HookHandler
}

const (
	ServiceInitHook Hook = iota + 1
	CmdCreateDbHook
	CmdMigrateDbHook
)

var (
	regMap         registryMap
	runtimeKey     = regKey("runtime")
	serviceKey     = regKey("service")
	appKey         = regKey("app")
	grpcSessionKey = regKey("grpc-session")
	hooksKey       = regKey("hooks")
)

func init() {
	regMap = make(registryMap)
	regMap[runtimeKey] = make([]interface{}, 1)
	regMap[serviceKey] = make([]interface{}, 0)
	regMap[grpcSessionKey] = make([]interface{}, 1)
	regMap[hooksKey] = make([]interface{}, 0)
	regMap[appKey] = make([]interface{}, 1)
}

func GlobalRegistry() GlobalRegistryInterface {
	return regMap
}

type GlobalRegistryInterface interface {
	Runtime() rt.Runtime
	SetRuntime(rt.Runtime)
	App() pkg.App
	SetApp(pkg.App)
	RegisterHook(Hook, HookHandler)
	InvokeHooks(Hook, ...interface{}) error
	ServiceLocator(string) ServiceLocator
}

func (m registryMap) App() pkg.App {
	return m[appKey][0].(pkg.App)
}

func (m registryMap) SetApp(app pkg.App) {
	m[appKey][0] = app
}

func (m registryMap) Runtime() rt.Runtime {
	return m[runtimeKey][0].(rt.Runtime)
}

func (m registryMap) SetRuntime(r rt.Runtime) {
	m[runtimeKey][0] = r
}

func (m registryMap) RegisterHook(typ Hook, handler HookHandler) {
	m[hooksKey] = append(m[hooksKey], &hookEntry{typ, handler})
}

func (m registryMap) InvokeHooks(typ Hook, args ...interface{}) error {
	for _, entry := range m[hooksKey] {
		entry := entry.(*hookEntry)
		if entry.typ != typ {
			continue
		}
		err := entry.handler(args...)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m registryMap) ServiceLocator(svcName string) ServiceLocator {
	return nil
}
