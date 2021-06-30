package server

import (
	"context"
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/service/hub/grpc/server/chello"
	//"github.com/AgentCoop/peppermint/internal/service/hub"
)

func (s *hubServer) Join(ctx context.Context, origReq *msg.Join_Request) (*msg.Join_Response, error) {
	desc := ctx.(grpc.ServerDescriptor)
	req := chello.NewJoin(origReq)
	desc.SetRequestData(req)

	sess := desc.Session()
	sess.Ipc().Grpc_Send(1, desc)
	v, ok := sess.Ipc().Grpc_Recv(1).(error)
	if ok {
		return nil, v
	}

	res := desc.ResponseData()
	return res.ToGrpc().(*msg.Join_Response), nil
}
