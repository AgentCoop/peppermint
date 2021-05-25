package server

import (
	"context"
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	srv "github.com/AgentCoop/peppermint/internal/grpc/server"
	context2 "github.com/AgentCoop/peppermint/internal/service/hub/grpc/context"
	"github.com/AgentCoop/peppermint/internal/service/hub/grpc/data/server/join"
	"github.com/AgentCoop/peppermint/internal/utils"
	//"github.com/AgentCoop/peppermint/internal/service/hub"
)

func (s *hubServer) Join(ctx context.Context, r *msg.Join_Request) (*msg.Join_Response, error) {
	//req := pair.(server.RequestResponsePair).GetRequest()
	pair := ctx.(srv.RequestResponsePair)
	req := join.NewJoin(pair, r)
	id := req.SessionId()

	desc, err := utils.GetSessDescriptorById(id)
	if err != nil  {
		return nil, err
	}

	joinJob := desc.Job()
	joinCtx := joinJob.GetValue().(context2.JoinContext)
	joinCtx.ReqChan()[1] <- pair

	return &msg.Join_Response{NodeId: 1}, nil
}
