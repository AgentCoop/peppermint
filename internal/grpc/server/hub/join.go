package hub

import (
	hub "github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	"context"
)

func (s *hubGrpcServer) Join(ctx context.Context, req *hub.Join_Request) (*hub.Join_Response, error) {
	return nil, nil
}
