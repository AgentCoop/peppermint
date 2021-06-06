package server

import (
	"context"
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/server/join"
)

func (s *hubServer) JoinHello(ctx context.Context, originalReq *msg.JoinHello_Request) (*msg.JoinHello_Response, error) {
	sess := join.CreateSession()
	callDesc := ctx.(grpc.ServerCallDesc)

	data := join.NewJoinHello(originalReq)
	callDesc.SetRequestData(data)

	sess.Ipc().Grpc_Send(0, callDesc)
	v, ok := sess.Ipc().Grpc_Recv(0).(error)
	if ok { return nil, v }

	callDesc.SetSessionId(sess.Id())
	res := callDesc.ResponseData()
	return res.ToGrpc().(*msg.JoinHello_Response), nil
}
