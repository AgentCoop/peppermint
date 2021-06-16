package client

import (
	"context"
	proto "github.com/AgentCoop/peppermint/internal/api/peppermint"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/calldesc"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc"
)

func prepareCallDescriptor(ctx context.Context, methodName string, svcPolicy runtime.ServicePolicy) g.ClientDescriptor {
	rt := runtime.GlobalRegistry().Runtime()
	switch v := ctx.(type) {
	case g.ClientDescriptor:
		return v
	default:
		cfg := rt.NodeConfigurator()
		method, _ := svcPolicy.FindMethodByName(methodName)
		var useEnc bool
		switch {
		case method.WasSet(proto.E_MEnforceEnc):
			useEnc = method.CallPolicy().EnforceEncryption()
		case svcPolicy.WasSet(proto.E_EnforceEnc):
			useEnc = svcPolicy.EnforceEncryption()
		default:
			useEnc = cfg.E2E_EncryptionEnabled()
		}
		secPolicy := calldesc.NewSecurityPolicy(useEnc, cfg.EncKey())
		callDesc := calldesc.NewClient(ctx, secPolicy, method.CallPolicy())
		callDesc.HandleMeta()
		return callDesc
	}
}

func PreUnaryInterceptor(svcPolicy runtime.ServicePolicy) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		callDesc := prepareCallDescriptor(ctx, method, svcPolicy)
		callDesc.Meta().SendHeader(nil)
		err := invoker(callDesc, method, req, reply, cc, opts...)
		return err
	}
}
