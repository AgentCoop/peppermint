package calldesc

import (
	"context"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/deps"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func NewSecurityPolicy(useEnc bool, encKey []byte) secPolicy {
	p := secPolicy{
		encKey:  encKey,
		e2e_Enc: useEnc,
	}
	return p
}

func NewServer(ctx context.Context, cfg deps.Configurator, secPolicy secPolicy) *sCallDesc {
	stream := grpc.ServerTransportStreamFromContext(ctx)
	desc := &sCallDesc{}
	desc.Context = ctx
	desc.common.typ = ServerCallDesc
	desc.meta.parent = &desc.common
	desc.secPolicy = secPolicy
	header, _ := metadata.FromIncomingContext(ctx)
	desc.meta.header = header
	desc.method = stream.Method()
	return desc
}

func NewServerInsecure(ctx context.Context, cfg deps.Configurator) *sCallDesc {
	secPolicy := secPolicy{e2e_Enc: false}
	desc := NewServer(ctx, cfg, secPolicy)
	return desc
}

func NewClientInSecure(ctx context.Context) *cCallDesc {
	secPolicy := secPolicy{e2e_Enc: false}
	desc := NewClient(ctx, secPolicy)
	return desc
}

func NewClient(ctx context.Context, secPolicy secPolicy) *cCallDesc {
	desc := &cCallDesc{}
	desc.common.Context = ctx
	desc.common.typ = ClientCallDesc
	desc.meta.parent = &desc.common
	desc.meta.header = metadata.New(nil)
	desc.secPolicy = secPolicy
	// Assign value to the node ID header
	rt := runtime.GlobalRegistry().Runtime()
	cfg := rt.NodeConfigurator()
	g.SetNodeId(&desc.meta.header, cfg.ExternalId())
	return desc
}
