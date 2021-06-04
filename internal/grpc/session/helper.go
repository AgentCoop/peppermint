package session

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Ipc_CallDesc(j job.Job, chanIdx int) (server.GrpcCallDescriptor, grpc.GrpcServiceLayersIpc) {
	v := j.GetValue()
	s, ok := v.(grpc.Session)
	if !ok {
		panic(status.Error(codes.Internal, "sys: job value must be a session object"))
	}
	ipc := s.Ipc()
	v = ipc.Svc_Recv(chanIdx)
	if v == nil {
		panic(status.Error(codes.Internal, "sys: expected call descriptor, got nil"))
	}
	callDesc, ok := v.(server.GrpcCallDescriptor)
	if !ok {
		panic(status.Error(codes.Internal, fmt.Sprintf("sys: expected call descriptor, got %v", v)))
	}
	return callDesc, ipc
}
