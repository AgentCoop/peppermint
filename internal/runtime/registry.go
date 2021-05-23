package runtime

import (
	"github.com/AgentCoop/peppermint/internal/service"
)

type parserCmdHook func(interface{})
type regKey string
type registryMap map[regKey][]interface{}

var (
	regMap registryMap
	runtimeKey  = regKey("runtime")
	serviceKey = regKey("service")
	parserCmdHookKey = regKey("parser-hook")
)

func init() {
	regMap = make(registryMap)
	regMap[runtimeKey] = make([]interface{}, 1)
	regMap[serviceKey] = make([]interface{}, 0)
	regMap[parserCmdHookKey] = make([]interface{}, 0)
}

func GlobalRegistry() GlobalRegistryInterface {
	return regMap
}

type GlobalRegistryInterface interface {
	Runtime() Runtime
	SetRuntime(Runtime)
	RegisterService(string, service.Service)
	RegisterParserCmdHook(string, parserCmdHook)
	LookupParserCmdHook(string) []parserCmdHook
}

func (m registryMap) Runtime() Runtime {
	return m[runtimeKey][0].(Runtime)
}

func (m registryMap) SetRuntime(r Runtime) {
	m[runtimeKey][0] = r
}

type serviceDesc struct {
	name string
	service service.Service
}

func (m registryMap) RegisterService(name string, service service.Service) {
	m[serviceKey] = append(m[serviceKey], &serviceDesc{name, service})
}

type parserCmdHookDesc struct {
	cmdName string
	hook parserCmdHook
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