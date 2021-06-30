//
// The main mechanism of communication between the gRPC and service layer.
//
package grpc

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
)

type Session interface {
	Id() i.SessionId
	Ipc() GrpcServiceLayersIpc
	Job() job.Job
	TaskContext() interface{}
	WithTaskContext(interface{})
}

// ********************************************** Server-side **********************************************************
// [Unary call, support for client pseudo-streaming]
// Service Layer										gRPC Layer
//----------------------------------------------------------------------------------------------------------------------
//														err := Grpc_Send(chanIdx, clientMsg)
// clientMsg|nil := Svc_Recv(chanIdx)
// Svc_Send(chanIdx, srvMsg)
//														srvMsg|err := Grpc_Recv(chanIdx)
// t.Done() for an unary call
// t.Tick() for client pseudo-streaming calls

// [Streaming]
// Service Layer										gRPC Layer
//----------------------------------------------------------------------------------------------------------------------
//														Grpc_MakeStreamable(chanIdx, stream)
//														Grpc_Send(chanIdx, clientMsg)
// clientMsg := Svc_Recv(chanIdx)
// *Svc_StreamSend(chanIdx, srvMsg)
// Svc_StreamClose(chanIdx)
//														err|nil := Grpc_WaitForStreamClose(chanIdx)

type GrpcServiceLayersIpc interface {
	Svc_Recv(chanIdx int) interface{}
	Svc_Send(chanIdx int, data interface{})
	Svc_StreamSend(chanIdx uint, data interface{})
	Svc_StreamClose(chanIdx uint)

	Grpc_Recv(chanIdx int) interface{}
	Grpc_Send(chanIdx int, data interface{})
	Grpc_MakeStreamable(chanIdx uint, stream Stream)
	Grpc_SendStreamable(chanId int, data interface{}) // For client pseudo-streaming
	Grpc_WaitForStreamClose(chanIdx uint) error
	Job() job.Job
}
