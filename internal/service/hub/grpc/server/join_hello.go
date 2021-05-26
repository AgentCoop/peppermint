package server

import (
	"context"
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	srv "github.com/AgentCoop/peppermint/internal/grpc/server"
	data "github.com/AgentCoop/peppermint/internal/service/hub/grpc/data/server/join"
	"github.com/AgentCoop/peppermint/internal/service/hub/service/server/join"
	"time"
)

func (s *hubServer) JoinHello(ctx context.Context, originalReq *msg.JoinHello_Request) (*msg.JoinHello_Response, error) {
	callDesc := ctx.(srv.GrpcCallDescriptor)
	req := data.NewJoinHello(callDesc, originalReq)
	_ = req.Validate()

	comm := srv.NewCommunicator(time.Minute)
	joinCtx := join.NewJoinContext()
	j := comm.Job()
	j.AddTask(joinCtx.JoinHelloTask)
	j.AddTask(joinCtx.JoinTask)
	j.Run()

	// Dispatch request-response pair to JoinHelloTask started above
	comm.GrpcTx(0, callDesc)
	v := comm.GrpcRx(0)

	switch v.(type) {
	case error:
		return nil, v.(error)
	default:
		res := callDesc.GetResponse()
		res.SetSessionId(comm.SessionId())
		return res.ToGrpcResponse().(*msg.JoinHello_Response), nil
	}
}
