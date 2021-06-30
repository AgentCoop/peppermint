package calldesc

import (
	"context"
	proto "github.com/AgentCoop/peppermint/internal/api/peppermint"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/pkg/node"
	"github.com/AgentCoop/peppermint/pkg/service"
	"google.golang.org/grpc/metadata"
)

func NewSecurityPolicy(useEnc bool, encKey []byte) *secPolicy {
	p := &secPolicy{
		encKey:  encKey,
		e2e_Enc: useEnc,
	}
	return p
}

func NewSecurityPolicyFromMethod(method service.Method, cfg node.NodeConfigurator) *secPolicy {
	var useEnc bool
	var encKey []byte
	switch {
	case method.WasSet(proto.E_MEnforceEnc):
		useEnc = method.CallPolicy().EnforceEncryption()
		encKey = cfg.EncKey()
	case method.ServicePolicy().WasSet(proto.E_EnforceEnc):
		useEnc = method.ServicePolicy().EnforceEncryption()
		encKey = cfg.EncKey()
	default:
		useEnc = cfg.E2E_EncryptionEnabled()
		encKey = cfg.EncKey()
	}
	secPolicy := NewSecurityPolicy(useEnc, encKey)
	return secPolicy
}

func NewServer(ctx context.Context, cfg service.ServiceConfigurator, method service.Method, secPolicy *secPolicy) *srvDescriptor {
	desc := &srvDescriptor{}
	desc.Context = ctx
	desc.common.typ = ServerType
	desc.meta.parent = &desc.common
	desc.secPolicy = secPolicy
	header, _ := metadata.FromIncomingContext(ctx)
	desc.meta.header = header
	desc.method = method
	desc.svcCfg = cfg
	return desc
}

func NewClient(ctx context.Context, secPolicy *secPolicy, method service.Method) *clDescriptor {
	desc := &clDescriptor{}
	desc.common.Context = ctx
	desc.common.typ = ClientType
	desc.common.method = method
	desc.meta.parent = &desc.common
	desc.meta.header = metadata.New(nil)
	desc.secPolicy = secPolicy
	// Assign value to the node ID header
	rt := runtime.GlobalRegistry().Runtime()
	cfg := rt.NodeConfigurator()
	g.SetNodeId(&desc.meta.header, cfg.ExternalId())
	desc.meta.nodeId = cfg.ExternalId()
	return desc
}
