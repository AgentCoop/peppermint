package hub

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"context"
)

func (s *server) Join(context.Context, *hub.Join_Request) (*hub.Join_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
