package server

import (
	"context"
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	srv "github.com/AgentCoop/peppermint/internal/grpc/server"
	"github.com/AgentCoop/peppermint/internal/grpc/session"
	"github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/server/join"
	//"github.com/AgentCoop/peppermint/internal/service/hub"
)

func (s *hubServer) Join(ctx context.Context, r *msg.Join_Request) (*msg.Join_Response, error) {
	callDesc := ctx.(srv.GrpcCallDescriptor)
	req := join.NewJoin(callDesc, r)
	id := req.SessionId()
	desc, _ := session.DescriptorById(id)
	_ = desc
	return nil, nil
	//if err != nil  { return nil, err }
	//
	//joinJob := desc.Job()
	//comm := joinJob.GetValue().(runtime.GrpcServiceCommunicator)
	//comm.GrpcTx(1, callDesc)
	//v := comm.GrpcRx(1)
	//
	//switch v.(type) {
	//case error:
	//	return nil, v.(error)
	//default:
	//	res := callDesc.GetResponse()
	//	res.SetSessionId(comm.SessionId())
	//	return res.ToGrpcResponse().(*msg.Join_Response), nil
	//}
}
