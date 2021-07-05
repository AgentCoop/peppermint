package calldesc

import (
	"context"
	proto "github.com/AgentCoop/peppermint/internal/api/peppermint"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/pkg"
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

func NewSecurityPolicyFromMethod(method service.Method, node pkg.Node) *secPolicy {
	var useEnc bool
	var encKey []byte
	switch {
	case method.WasSet(proto.E_MEnforceEnc):
		useEnc = method.CallPolicy().EnforceEncryption()
		encKey = node.EncKey()
	case method.ServicePolicy().WasSet(proto.E_EnforceEnc):
		useEnc = method.ServicePolicy().EnforceEncryption()
		encKey = node.EncKey()
	default:
		useEnc = node.EncEnabled()
		encKey = node.EncKey()
	}
	secPolicy := NewSecurityPolicy(useEnc, encKey)
	return secPolicy
}

func NewServer(ctx context.Context, svc service.Service, method service.Method, secPolicy *secPolicy) *srvDescriptor {
	desc := &srvDescriptor{}
	desc.Context = ctx
	desc.common.typ = ServerType
	desc.meta.parent = &desc.common
	desc.secPolicy = secPolicy
	header, _ := metadata.FromIncomingContext(ctx)
	desc.meta.header = header
	desc.method = method
	desc.svc = svc
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
	app := runtime.GlobalRegistry().App().(pkg.AppNode)
	node := app.Node()
	g.SetNodeId(&desc.meta.header, node.ExternalId())
	desc.meta.nodeId = node.ExternalId()
	return desc
}
