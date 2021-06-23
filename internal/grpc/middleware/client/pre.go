package client

import (
	"context"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/stream"
	"github.com/AgentCoop/peppermint/internal/grpc/calldesc"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc"
)

func prepareCallDescriptor(ctx context.Context, c g.BaseClient, methodName string, svcPolicy runtime.ServicePolicy) g.ClientDescriptor {
	rt := runtime.GlobalRegistry().Runtime()
	switch v := ctx.(type) {
	case g.ClientDescriptor:
		return v
	default:
		cfg := rt.NodeConfigurator()
		method, _ := svcPolicy.FindMethodByName(methodName)
		secPolicy := calldesc.NewSecurityPolicyFromMethod(method, cfg)
		desc := calldesc.NewClient(ctx, secPolicy, method.CallPolicy())
		if method.CallPolicy().SessionSticky() {
			lastCall := c.LastCall()
			if lastCall == nil { panic("lastCall nil") }
			desc.WithSessionFrom(lastCall)
		}
		//desc.HandleMeta()
		return desc
	}
}

func PreUnaryInterceptor(client g.BaseClient, svcPolicy runtime.ServicePolicy) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		desc := prepareCallDescriptor(ctx, client, method, svcPolicy)
		desc.Meta().SendHeader(nil)
		cs, _ := cc.NewStream(ctx, nil, method, opts...)
		stream.NewClientStream(cs, desc)
		err := invoker(desc, method, req, reply, cc, opts...)
		return err
	}
}
