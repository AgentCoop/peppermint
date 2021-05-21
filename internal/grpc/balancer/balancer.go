package balancer

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type upstream struct {

}

type downstream struct {
	ss *grpc.ServerStream
	cs *grpc.ClientStream
}

type balancer struct {
	server.BaseServer
	handler *grpc.StreamHandler
}


func (b *balancer) unknownStreamHandler(srv interface{}, stream grpc.ServerStream) error {
	fullMethod, _ := grpc.MethodFromServerStream(stream)
	balancerJob := newBalancerJob(fullMethod, stream)
	<-balancerJob.Run()
	// @todo must return error as second argument instead of interface{}
	_, _ = balancerJob.GetInterruptedBy()
	return nil
}

func NewGrpcBalancer(address string) *balancer {
	b := new(balancer)
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
