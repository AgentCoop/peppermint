package calldesc

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/grpc/client"
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

func NewClientCallDesc(ctx context.Context, client client.BaseClient, secPolicy secPolicy) cCallDesc {
	desc := cCallDesc{Context: ctx}
	desc.client = client
	desc.secPolicy = secPolicy
	return desc
}

func NewServerCallDesc(ctx context.Context, cfg deps.Configurator, secPolicy secPolicy) *sCallDesc {
	desc := &sCallDesc{Context: ctx}
	desc.secPolicy = secPolicy
	header, _ := metadata.FromIncomingContext(ctx)
	desc.meta.header = header
	return desc
}