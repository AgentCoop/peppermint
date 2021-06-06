package proxy

import (
	job "github.com/AgentCoop/go-work"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	s "github.com/AgentCoop/peppermint/internal/grpc/stream"
	"github.com/AgentCoop/peppermint/internal/grpc/client"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"google.golang.org/grpc"
)

type proxyConn struct {
	upClient       client.BaseClient
	downstream     g.ServerStream
	upstream       g.ClientStream
	downstreamChan chan codec.Packet
	upstreamChan   chan codec.Packet
}

func NewProxyConnJob(upstreamClient client.BaseClient, downstream grpc.ServerStream, downKey []byte) (job.Job, error) {
	//fullMethodName, ok := grpc.MethodFromServerStream(downstream)
	//if !ok {
	//	return nil, status.Error(codes.InvalidArgument, "failed to retrieve gRPC method name")
	//}
	down := s.NewServerStream(downstream, nil)
	pconn := &proxyConn{
		upClient:       upstreamClient,
		downstream:     down,
		downstreamChan: make(chan codec.Packet, 1),
		upstreamChan:   make(chan codec.Packet, 1),
	}
	pconnJob := job.NewJob(nil)
	pconnJob.AddOneshotTask(upstreamClient.ConnectTask)
	pconnJob.AddTask(pconn.initTask)
	pconnJob.AddTask(pconn.readUpstreamTask)
	pconnJob.AddTask(pconn.readDownstreamTask)
	pconnJob.AddTask(pconn.writeStreamTask)
	return pconnJob, nil
}
