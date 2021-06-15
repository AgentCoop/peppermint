package client

import (
	"context"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/calldesc"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc"
)

func prepareCallDescriptor(ctx context.Context, methodName string, svcPolicy g.ServicePolicy) g.ClientCallDesc {
	rt := runtime.GlobalRegistry().Runtime()
	switch v := ctx.(type) {
	case g.ClientCallDesc:
		return v
	default:
		cfg := rt.NodeConfigurator()
		method, _ := svcPolicy.FindMethodByName(methodName)
		methodPolicy := method.CallPolicy()
		useEnc :=  methodPolicy.IsSecure() || cfg.E2E_EncryptionEnabled()
		secPolicy := calldesc.NewSecurityPolicy(useEnc, cfg.EncKey())
		callDesc := calldesc.NewClient(ctx, secPolicy, methodPolicy)
		callDesc.HandleMeta()
		return callDesc
	}
}

func PreUnaryInterceptor(svcPolicy g.ServicePolicy) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		callDesc := prepareCallDescriptor(ctx, method, svcPolicy)
		callDesc.Meta().SendHeader(nil)
		err := invoker(callDesc, method, req, reply, cc, opts...)
		return err
	}
}
