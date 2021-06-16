package session

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Ipc_CallDesc(j job.Job, chanIdx int) (grpc.ServerDescriptor, grpc.GrpcServiceLayersIpc) {
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
	callDesc, ok := v.(grpc.ServerDescriptor)
	if !ok {
		panic(status.Errorf(codes.Internal, "sys: expected call descriptor, got %v", v))
	}
	return callDesc, ipc
}

func FindById(id i.SessionId) (*sessionDesc, error) {
	desc, ok := sMap[id]
	if !ok  {
		return nil, status.Error(codes.DeadlineExceeded, "gRPC session has been expired or session ID is invalid")
	} else {
		return desc, nil
	}
}