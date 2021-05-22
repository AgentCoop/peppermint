package balancer

import (
	"context"
	job "github.com/AgentCoop/go-work"
	c "github.com/AgentCoop/peppermint/internal/grpc/client"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

type grpcServer struct {
	server.BaseServer
	handler *grpc.StreamHandler
}

type balancerCtx struct {
	c.BaseClient
	resolver *resolver
	upstreamCtx context.Context
	upstreamCancel context.CancelFunc
	stream grpc.ServerStream
}

type resolver struct {
	server.Session
	server.Request
	target net.Addr
}

func newResolver(stream grpc.ServerStream) *resolver {
	r := new(resolver)
	return r
}

func (r *resolver) bySessionId() {
	switch {
	case ctx.Request.SessionId() != 0: 	// Get target node by session id
		node, err := ctx.Session.Node(ctx.Request.SessionId())
		if err != nil { return nil, err }
		ctx.target = node.ServiceAddressByName("")
	default: // Select a target using balancer specified strategy
	}
}

func (r *resolver) run() net.Addr {

}

func NewBalancerContext(stream grpc.ServerStream) (*balancerCtx, error) {
	ctx := new(balancerCtx)
	ctx.stream = stream
	ctx.resolver.run()
	ctx.BaseClient = c.NewBaseClientWithContext(ctx.resolver.target, ctx.upstreamCtx)
	return ctx, nil
}

func (b *balancerCtx) setupUpstreamHeader() {
	md, _ := metadata.FromIncomingContext(b.stream.Context())
	b.upstreamCtx = metadata.NewOutgoingContext(b.upstreamCtx, md.Copy())
}

func (b *grpcServer) unknownStreamHandler(srv interface{}, stream grpc.ServerStream) error {
	fullMethod, _ := grpc.MethodFromServerStream(stream)
	ctx, err := NewBalancerContext(stream)
	if err != nil { return err }

	balancerJob := newBalancerJob(fullMethod, stream)
	balancerJob.AddOneshotTask(ctx.BaseClient.ConnectTask)
	balancerJob.AddTask(ctx.proxyCallTask)
	<-balancerJob.Run()
	// @todo must return error as second argument instead of interface{}
	_, _ = balancerJob.GetInterruptedBy()
	return nil
}

func NewGrpcBalancer(address string) *grpcServer {
	b := new(grpcServer)
	grpc := grpc.NewServer(
		grpc.UnknownServiceHandler(b.unknownStreamHandler),
	)
	b.BaseServer = server.NewBaseServer(address, grpc)
	return b
}

func newBalancerJob(serviceName string, stream grpc.ServerStream) job.Job {
	ctx := stream.Context()
	req := server.NewRequest(ctx)
	// If gRPC session is active, forward the call based on its value
	sId := req.SessionId()

	md, ok := metadata.FromIncomingContext(ctx)
	// Copy the inbound metadata explicitly.
	outCtx, _ := context.WithCancel(ctx)
	outCtx = metadata.NewOutgoingContext(outCtx, md.Copy())
}
