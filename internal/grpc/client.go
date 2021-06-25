package grpc

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	"google.golang.org/grpc"
	"context"
)

type ConnProvider func(grpc.ClientConnInterface)

type BaseClient interface {
	Context() context.Context
	ConnectTask(j job.Job) (job.Init, job.Run, job.Finalize)
	Connection() grpc.ClientConnInterface
	WithContext(ctx context.Context)
	WithConnProvider(ConnProvider)
	WithEncKey([]byte)
	WithUnaryInterceptors(...grpc.UnaryClientInterceptor)
	NodeId() i.NodeId
	IsSecure() bool
	EncKey() []byte
	SessionId() i.SessionId
	SetSessionId(id i.SessionId)

	WithTargetPort(uint16)
	WithTimeout(uint)

	LastCall() ClientDescriptor
	SetLastCall(ClientDescriptor)
}
