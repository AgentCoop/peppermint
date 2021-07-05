package runtime

import (
	"github.com/AgentCoop/peppermint/pkg"
	rt "github.com/AgentCoop/peppermint/pkg/runtime"
)

type parserCmdHook func(interface{})
type regKey string
type registryMap map[regKey][]interface{}

type Hook int
type HookHandler func(...interface{})
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
	regMap           registryMap
	runtimeKey       = regKey("runtime")
	serviceKey       = regKey("service")
	parserCmdHookKey = regKey("cli-parser-hook")
	appKey            = regKey("app")
	grpcSessionKey   = regKey("grpc-session")
	hooksKey         = regKey("hooks")
)

func init() {
	regMap = make(registryMap)
	regMap[runtimeKey] = make([]interface{}, 1)
	regMap[serviceKey] = make([]interface{}, 0)
	regMap[parserCmdHookKey] = make([]interface{}, 0)
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
	InvokeHooks(Hook, ...interface{})
	ServiceLocator(string) ServiceLocator
	RegisterParserCmdHook(string, parserCmdHook)
	LookupParserCmdHook(string) []parserCmdHook
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

func (m registryMap) InvokeHooks(typ Hook, args ...interface{}) {
	for _, entry := range m[hooksKey] {
		entry := entry.(*hookEntry)
		if entry.typ != typ { continue }
		entry.handler(args...)
	}
}

func (m registryMap) ServiceLocator(svcName string) ServiceLocator {
	return nil
}

type parserCmdHookDesc struct {
	cmdName string
	hook    parserCmdHook
}

func (m registryMap) RegisterParserCmdHook(cmdName string, hook parserCmdHook) {
	m[parserCmdHookKey] = append(m[parserCmdHookKey], &parserCmdHookDesc{cmdName, hook})
}

func (m registryMap) LookupParserCmdHook(cmdName string) []parserCmdHook {
	var out []parserCmdHook
	for _, v := range m[parserCmdHookKey] {
		vv := v.(*parserCmdHookDesc)
		if vv.cmdName == cmdName {
			out = append(out, vv.hook)
		}
	}
	return out
}
