package service

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint"
	"github.com/AgentCoop/peppermint/internal/grpc/protobuf"
)

type svcOptions struct {
	enforceEnc          bool
	defaultPort         int32
	ipcUnixDomainSocket string
}

type methodsMap map[string]*methodOptions
type methodOptions struct {
	enforceEnc    bool
	streamable    bool
	newSession    int32
	requiredRoles []string
}

type svcPolicy struct {
	svcFullName string
	desc        protobuf.ServiceDescriptor
	sOpts       svcOptions
	methods     methodsMap
}

type method struct {
	name string
	opts *methodOptions
}

func (p *svcPolicy) populate(methods []string) {
	svcOptions := protobuf.NewSvcLevelOptions()
	svcOptions.AddItem(peppermint.E_EnforceEnc, &p.sOpts.enforceEnc)
	svcOptions.AddItem(peppermint.E_Port, &p.sOpts.defaultPort)

	sm := protobuf.NewMethodLevelOptions(methods)
	for methodName, _ := range sm {
		mOpt := &methodOptions{}
		p.methods[methodName] = mOpt
		sm.AddItem(methodName, peppermint.E_MEnforceEnc, &mOpt.enforceEnc)
		sm.AddItem(methodName, peppermint.E_Streamable, &mOpt.streamable)
		sm.AddItem(methodName, peppermint.E_NewSession, &mOpt.newSession)
	}

	p.desc.FetchServiceCustomOptions(svcOptions, sm)
	sm.OverrideVal(svcOptions[peppermint.E_EnforceEnc], peppermint.E_MEnforceEnc)
}

func (p *svcPolicy) EnforceEncryption() bool {
	return p.sOpts.enforceEnc
}

func (p *svcPolicy) DefaultPort() int {
	return int(p.sOpts.defaultPort)
}

func (p *svcPolicy) Ipc_UnixDomainSocket() string {
	return p.sOpts.ipcUnixDomainSocket
}

func (p *svcPolicy) FindMethodByName(shortName string) (method, bool) {
	m := method{}
	for methodName, opts := range p.methods {
		if methodName == shortName {
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

func (m method) CallPolicy() *methodOptions {
	return m.opts
}

func (m *methodOptions) IsSecure() bool {
	return false
}

func (m *methodOptions) IsStreamable() bool {
	return m.streamable
}

func (m *methodOptions) OpenNewSession() int {
	return int(m.newSession)
}

func (m *methodOptions) RequiredRoles() []string {
	return m.requiredRoles
}
