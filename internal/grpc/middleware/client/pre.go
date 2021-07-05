package client

import (
	"context"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/calldesc"
	"github.com/AgentCoop/peppermint/internal/grpc/stream"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/pkg"
	"github.com/AgentCoop/peppermint/pkg/service"
	"google.golang.org/grpc"
	"time"
)

func prepareCallDescriptor(ctx context.Context, cc g.BaseClient, methodName string, svcPolicy service.ServicePolicy) g.ClientDescriptor {
	app := runtime.GlobalRegistry().App().(pkg.AppNode)
	switch v := ctx.(type) {
	case g.ClientDescriptor:
		return v
	default:
		method, _ := svcPolicy.FindMethodByName(methodName)
		// Set call timeout in milliseconds if needed
		timeout := method.CallPolicy().Timeout()
		if timeout > 0 {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, time.Duration(timeout)*time.Millisecond)
			_ = cancel
		}
		secPolicy := calldesc.NewSecurityPolicyFromMethod(method, app.Node())
		desc := calldesc.NewClient(ctx, secPolicy, method)
		if method.CallPolicy().SessionSticky() {
			lastCall := cc.LastCall()
			if lastCall == nil {
				panic("lastCall nil")
			}
			desc.WithSessionFrom(lastCall)
		}
		//desc.HandleMeta()
		return desc
	}
}

func PreUnaryInterceptor(client g.BaseClient, svcPolicy service.ServicePolicy) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		desc := prepareCallDescriptor(ctx, client, method, svcPolicy)
		desc.Meta().SendHeader(nil)
		cs, _ := cc.NewStream(ctx, nil, method, opts...)
		stream.NewClientStream(cs, desc)
		err := invoker(desc, method, req, reply, cc, opts...)
		return err
	}
}
