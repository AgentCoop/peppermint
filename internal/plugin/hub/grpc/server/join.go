package server

import (
	"context"
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"

	//"github.com/AgentCoop/peppermint/internal/service/hub"
)

func (s *hubServer) Join(ctx context.Context, r *msg.Join_Request) (*msg.Join_Response, error) {
	return nil, nil
	//callDesc := ctx.(srv.GrpcCallDescriptor)
	//req := join.NewJoin(callDesc, r)
	//id := req.SessionId()
	//desc, err := grpc.GetSessDescriptorById(id)
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
