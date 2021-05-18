package hub

import (
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	"context"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
)

func (s *hubServer) Join(pair context.Context, msg *msg.Join_Request) (*msg.Join_Response, error) {

	req := pair.(server.RequestResponsePair).GetRequest()
	id := req.SessionId()
	_ = id
	//sessId := reqHeader.SessionId()
	//_ = sessId

	return nil, nil
}
