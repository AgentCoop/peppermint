package server

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/frontoffice/test"
)

func (t *testServer) Single(ctx context.Context, req *test.Request) (*test.Response, error) {
	return nil, nil
}
