package service

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint"
	"github.com/AgentCoop/peppermint/internal/grpc/protobuf"
)

type methodOptions struct {
	enableEnc, disableEnc bool
	streamable            bool
	newSession            int32
}

type serviceOptions struct {
	enforceEnc          bool
	defaultPort         int32
	ipcUnixDomainSocket string
}

type svcPolicy struct {
	desc            protobuf.ServiceDescriptor
	mOptsReceiver   map[string]*methodOptions // receiver for method option values
	svcOptsReceiver serviceOptions
}

func (p *svcPolicy) populate() {
	svcOptions := protobuf.ServiceLevelOptions{
		peppermint.E_EnforceEnc:    &p.svcOptsReceiver.enforceEnc,
		peppermint.E_IpcUnixSocket: &p.svcOptsReceiver.ipcUnixDomainSocket,
		peppermint.E_Port:          &p.svcOptsReceiver.defaultPort,
	}
	mOpts := protobuf.MethodLevelOptions{}
	for methodName, opts := range p.mOptsReceiver {
		mOpts.AddItem(methodName, peppermint.E_EnableEnc, &opts.enableEnc)
		mOpts.AddItem(methodName, peppermint.E_DisableEnc, &opts.disableEnc)
		mOpts.AddItem(methodName, peppermint.E_Streamable, &opts.streamable)
		mOpts.AddItem(methodName, peppermint.E_NewSession, &opts.newSession)
	}
	p.desc.FetchServiceCustomOptions(svcOptions, mOpts)
}

func (p *svcPolicy) EnforceEncryption() bool {
	return p.svcOptsReceiver.enforceEnc
}

func (p *svcPolicy) DefaultPort() int {
	return int(p.svcOptsReceiver.defaultPort)
}

func (p *svcPolicy) Ipc_UnixDomainSocket() string {
	return p.svcOptsReceiver.ipcUnixDomainSocket
}
