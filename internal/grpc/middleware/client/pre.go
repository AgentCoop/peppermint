package client

import (
	"context"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/calldesc"
	"github.com/AgentCoop/peppermint/internal/grpc/client"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc"
)

func prepareCallDescriptor(ctx context.Context, client client.BaseClient) g.ClientCallDesc {
	rt := runtime.GlobalRegistry().Runtime()
	switch v := ctx.(type) {
	case g.ClientCallDesc:
		return v
	default:
		cfg := rt.NodeConfigurator()
		secPolicy := calldesc.NewSecurityPolicy(cfg.E2E_EncryptionEnabled(), cfg.EncKey())
		callDesc := calldesc.NewClient(ctx, secPolicy)
		callDesc.HandleMeta()
		return callDesc
	}
}

func PreUnaryInterceptor(c client.BaseClient) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		callDesc := prepareCallDescriptor(ctx, c)
		callDesc.Meta().SendHeader(nil)
		err := invoker(callDesc, method, req, reply, cc, opts...)
		return err
	}
}
