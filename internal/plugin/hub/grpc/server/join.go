package server

import (
	"context"
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/session"
	"github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/server/join"
	//"github.com/AgentCoop/peppermint/internal/service/hub"
)

func (s *hubServer) Join(ctx context.Context, origReq *msg.Join_Request) (*msg.Join_Response, error) {
	desc := ctx.(grpc.ServerDescriptor)
	req := join.NewJoin(origReq)
	desc.SetRequestData(req)

	id := desc.Meta().SessionId()
	sess, err := session.FindById(id)
	if err != nil { return nil, err }

	sess.Ipc().Grpc_Send(1, desc)
	v, ok := sess.Ipc().Grpc_Recv(1).(error)
	if ok { return nil, v }

	res := desc.ResponseData()
	return res.ToGrpc().(*msg.Join_Response), nil
}
