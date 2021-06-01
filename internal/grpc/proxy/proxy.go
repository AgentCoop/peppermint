package proxy

import (
	job "github.com/AgentCoop/go-work"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/client"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type proxyConn struct {
	upClient       client.BaseClient
	downstream     runtime.Stream
	upstream       runtime.Stream
	downstreamChan chan codec.Packet
	upstreamChan   chan codec.Packet
}

func NewProxyConnJob(upstreamClient client.BaseClient, downstream grpc.ServerStream, downKey []byte) (job.Job, error) {
	fullMethodName, ok := grpc.MethodFromServerStream(downstream)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "failed to retrieve gRPC method name")
	}
	down := g.NewStream(downstream, fullMethodName, downKey)
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
