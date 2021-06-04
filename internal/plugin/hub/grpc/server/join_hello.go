package server

import (
	"context"
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	srv "github.com/AgentCoop/peppermint/internal/grpc/server"
	"github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/server/join"
)

func (s *hubServer) JoinHello(ctx context.Context, originalReq *msg.JoinHello_Request) (*msg.JoinHello_Response, error) {
	sess := join.CreateSession()
	callDesc := ctx.(srv.GrpcCallDescriptor)
	join.NewJoinHello(callDesc, originalReq)

	sess.Ipc().Grpc_Send(0, callDesc)
	v, ok := sess.Ipc().Grpc_Recv(0).(error)
	if ok { return nil, v }

	res := callDesc.GetResponse()
	res.SetSessionId(sess.Id())
	return res.ToGrpcResponse().(*msg.JoinHello_Response), nil
}
