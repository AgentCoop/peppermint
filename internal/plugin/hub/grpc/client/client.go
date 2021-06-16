package client

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/grpc"
)

type HubClient interface {
	grpc.BaseClient
	hub.HubClient
}

type hubClient struct {
	grpc.BaseClient
	hub.HubClient
}
