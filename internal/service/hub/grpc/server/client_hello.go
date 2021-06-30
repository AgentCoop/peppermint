package server

import (
	"context"
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/service/hub/grpc/server/chello"
)

func (s *hubServer) ClientHello(ctx context.Context, origReq *msg.ClientHello_Request) (*msg.ClientHello_Response, error) {
	desc := ctx.(grpc.ServerDescriptor)
	sess := desc.Session()

	cHelloCtx := chello.NewContext()
	sess.WithTaskContext(cHelloCtx)
	sess.Job().AddTask(cHelloCtx.ClientHelloTask)
	sess.Job().AddTask(cHelloCtx.JoinTask)
	sess.Job().Run()

	req := chello.NewClientHelloRequest(origReq)
	desc.SetRequestData(req)

	sess.Ipc().Grpc_Send(0, desc)
	v, ok := sess.Ipc().Grpc_Recv(0).(error)
	if ok { return nil, v }

	resp := desc.ResponseData()
	return resp.ToGrpc().(*msg.ClientHello_Response), nil
}
