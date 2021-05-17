package hub

import (
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	"context"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
)

func (s *hubServer) Join(ctx context.Context, req *msg.Join_Request) (*msg.Join_Response, error) {

	reqHeader := ctx.(server.RequestHeader)
	sessId := reqHeader.SessionId()
	_ = sessId

	return nil, nil
}
