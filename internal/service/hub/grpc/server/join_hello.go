package server

import (
	"context"
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/service/hub/grpc/server/join"
)

func (s *hubServer) JoinHello(ctx context.Context, origReq *msg.JoinHello_Request) (*msg.JoinHello_Response, error) {
	sess := join.CreateSession()
	desc := ctx.(grpc.ServerDescriptor)

	req := join.NewJoinHello(origReq)
	desc.SetRequestData(req)

	sess.Ipc().Grpc_Send(0, desc)
	v, ok := sess.Ipc().Grpc_Recv(0).(error)
	if ok { return nil, v }

	desc.Meta().SetSessionId(sess.Id())
	resp := desc.ResponseData()
	return resp.ToGrpc().(*msg.JoinHello_Response), nil
}
