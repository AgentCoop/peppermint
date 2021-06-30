package service

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint"
	"github.com/AgentCoop/peppermint/internal/grpc/protobuf"
	"github.com/AgentCoop/peppermint/internal/utils"
	"github.com/AgentCoop/peppermint/pkg/service"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type svcOptions struct {
	enforceEnc          bool
	defaultPort         uint32
	ipcUnixDomainSocket string
}

type methodsMap map[string]*methodOptions
type methodOptions struct {
	enforceEnc    bool
	streamable    bool
	sessSticky    bool
	newSession    uint32
	closeSession  bool
	timeoutMs     uint32
	requiredRoles []string
}

type svcPolicy struct {
	svcFullName string
	desc        protobuf.ServiceDescriptor
	sOptsMap    protobuf.SvcLevelOptionsMap
	mOptsMap    protobuf.MethodLevelOptionsMap
	sOpts       svcOptions
	methods     methodsMap
}

type method struct {
	name     string
	fullName string
	policy   *svcPolicy
	opts     *methodOptions
}

func (p *svcPolicy) populate(methods []string) {
	p.sOptsMap = protobuf.NewSvcLevelOptions()
	p.sOptsMap.AddItem(peppermint.E_EnforceEnc, &p.sOpts.enforceEnc)
	p.sOptsMap.AddItem(peppermint.E_Port, &p.sOpts.defaultPort)
	p.sOptsMap.AddItem(peppermint.E_IpcUnixSocket, &p.sOpts.ipcUnixDomainSocket)

	p.mOptsMap = protobuf.NewMethodLevelOptions(methods)
	for methodName, _ := range p.mOptsMap {
		mOpt := &methodOptions{}
		p.methods[methodName] = mOpt
		p.mOptsMap.AddItem(methodName, peppermint.E_MEnforceEnc, &mOpt.enforceEnc)
		p.mOptsMap.AddItem(methodName, peppermint.E_Streamable, &mOpt.streamable)
		p.mOptsMap.AddItem(methodName, peppermint.E_NewSession, &mOpt.newSession)
		p.mOptsMap.AddItem(methodName, peppermint.E_CloseSession, &mOpt.closeSession)
		p.mOptsMap.AddItem(methodName, peppermint.E_SessionSticky, &mOpt.sessSticky)
		p.mOptsMap.AddItem(methodName, peppermint.E_Timeout, &mOpt.timeoutMs)
	}
	p.desc.FetchServiceCustomOptions(p.sOptsMap)
	p.desc.FetchMethodsCustomOptions(p.mOptsMap)
}

func (p *svcPolicy) WasSet(ext *protoimpl.ExtensionInfo) bool {
	return p.sOptsMap[ext].WasSet()
}

func (p *svcPolicy) EnforceEncryption() bool {
	return p.sOpts.enforceEnc
}

func (p *svcPolicy) DefaultPort() uint16 {
	max := uint32(^uint16(0))
	if p.sOpts.defaultPort > max {
		panic("port number exceeds maximum allowed value of 65,535")
	}
	return uint16(p.sOpts.defaultPort)
}

func (p *svcPolicy) Ipc_UnixDomainSocket() string {
	return p.sOpts.ipcUnixDomainSocket
}

func (p *svcPolicy) FindMethodByName(name string) (service.Method, bool) {
	name = utils.Conv_FromLongToShortMethod(name)
	m := method{}
	m.policy = p
	for methodName, opts := range p.methods {
		if methodName == name {
			m.name = methodName
			m.opts = opts
			return m, true
		}
	}
	return m, false
}

func (m method) Name() string {
	return m.name
}

func (m method) FullName() string {
	return "/" + m.policy.svcFullName + "/" + m.name
}

func (m method) ServicePolicy() service.ServicePolicy {
	return m.policy
}

func (m method) CallPolicy() service.MethodCallPolicy {
	return m.opts
}

func (m method) WasSet(ext *protoimpl.ExtensionInfo) bool {
	opts := m.policy.mOptsMap[m.name]
	if _, has := opts[ext]; !has {
		return false
	}
	return opts[ext].WasSet()
}

func (m *methodOptions) EnforceEncryption() bool {
	return m.enforceEnc
}

func (m *methodOptions) IsStreamable() bool {
	return m.streamable
}

func (m *methodOptions) SessionSticky() bool {
	return m.sessSticky
}

func (m *methodOptions) NewSession() int {
	return int(m.newSession)
}

func (m *methodOptions) CloseSession() bool {
	return m.closeSession
}

func (m *methodOptions) Timeout() uint32 {
	return m.timeoutMs
}

func (m *methodOptions) RequiredRoles() []string {
	return m.requiredRoles
}
