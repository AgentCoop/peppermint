package calldesc

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/runtime/deps"
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
	desc := &sCallDesc{}
	desc.Context = ctx
	desc.common.typ = ServerCallDesc
	desc.meta.parent = &desc.common
	desc.secPolicy = secPolicy
	header, _ := metadata.FromIncomingContext(ctx)
	desc.meta.header = header
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
	return desc
}
