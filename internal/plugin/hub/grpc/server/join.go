package server

import (
	"context"
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/session"
	"github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/server/join"
	//"github.com/AgentCoop/peppermint/internal/service/hub"
)

func (s *hubServer) Join(ctx context.Context, r *msg.Join_Request) (*msg.Join_Response, error) {
	callDesc := ctx.(grpc.ServerCallDesc)
	req := join.NewJoin(r)
	callDesc.SetRequestData(req)

	id := callDesc.Meta().SessionId()
	sess, err := session.FindById(id)
	if err != nil { return nil, err }

	sess.Ipc().Grpc_Send(1, callDesc)
	v, ok := sess.Ipc().Grpc_Recv(1).(error)
	if ok { return nil, v }

	res := callDesc.ResponseData()
	return res.ToGrpc().(*msg.Join_Response), nil
}
