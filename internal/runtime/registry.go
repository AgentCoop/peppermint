package runtime

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/db"
	"time"
)

type parserCmdHook func(interface{})
type regKey string
type registryMap map[regKey][]interface{}

var (
	regMap registryMap
	runtimeKey  = regKey("runtime")
	serviceKey = regKey("service")
	parserCmdHookKey = regKey("cli-parser-hook")
	dbKey = regKey("db")
	grpcSessionKey = regKey("grpc-session")
)

func init() {
	regMap = make(registryMap)
	regMap[runtimeKey] = make([]interface{}, 1)
	regMap[serviceKey] = make([]interface{}, 0)
	regMap[parserCmdHookKey] = make([]interface{}, 0)
	regMap[dbKey] = make([]interface{}, 1)
	regMap[grpcSessionKey] = make([]interface{}, 1)
}

func GlobalRegistry() GlobalRegistryInterface {
	return regMap
}

type SessionDesc interface {
	Job() job.Job
	Expired() bool
}

type Session interface {
	New(job job.Job, expireInSecs time.Duration) i.SessionId
	Lookup(i.SessionId) SessionDesc
	Remove(i.SessionId)
}

type GlobalRegistryInterface interface {
	Runtime() Runtime
	SetRuntime(Runtime)

	Db() db.Db
	SetDb(db.Db)

	GrpcSession() Session
	SetGrpcSession(Session)

	RegisterService(*ServiceInfo)
	Services() []*ServiceInfo
	LookupService(string) *ServiceInfo

	RegisterParserCmdHook(string, parserCmdHook)
	LookupParserCmdHook(string) []parserCmdHook
}

func (m registryMap) Runtime() Runtime {
	return m[runtimeKey][0].(Runtime)
}

func (m registryMap) SetRuntime(r Runtime) {
	m[runtimeKey][0] = r
}

func (m registryMap) Db() db.Db {
	return m[dbKey][0].(db.Db)
}

func (m registryMap) SetDb(db db.Db) {
	m[dbKey][0] = db
}

func (m registryMap) GrpcSession() Session {
	return m[grpcSessionKey][0].(Session)
}

func (m registryMap) SetGrpcSession(s Session) {
	m[grpcSessionKey][0] = s
}

func (m registryMap) RegisterService(info *ServiceInfo) {
	m[serviceKey] = append(m[serviceKey], info)
}

func (m registryMap) Services() []*ServiceInfo {
	var out []*ServiceInfo
	out = make([]*ServiceInfo, 0)
	for _, v := range m[serviceKey] {
		vv := v.(*ServiceInfo)
		out = append(out, vv)
	}
	return out
}

func (m registryMap) LookupService(name string) *ServiceInfo {
	for _, v := range m[serviceKey] {
		vv := v.(*ServiceInfo)
		if vv.Name == name {
			return vv
		}
	}
	return nil
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