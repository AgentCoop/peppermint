package client

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/frontoffice/test"
	"github.com/AgentCoop/peppermint/internal/grpc"
)

type TestClient interface {
	grpc.BaseClient
	test.TestClient
}

type testClient struct {
	grpc.BaseClient
	test.TestClient
}

