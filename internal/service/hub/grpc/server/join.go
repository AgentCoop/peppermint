package server

import (
	"context"
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"github.com/AgentCoop/peppermint/internal/runtime"
	//"github.com/AgentCoop/peppermint/internal/service/hub"
)

func (s *hubServer) Join(pair context.Context, msg *msg.Join_Request) (*msg.Join_Response, error) {
	req := pair.(server.RequestResponsePair).GetRequest()
	id := req.SessionId()

	sess := runtime.GlobalRegistry().GrpcSession()
	desc := sess.Lookup(id)
	if desc == nil {
		return nil, nil
	}

	//joinJob := desc.Job()
	//joinCtx := joinJob.GetValue().(hub.JoinContext)

	//joinCtx.ReqChan()[1] <- nil

	return nil, nil
}
