package client

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	c "github.com/AgentCoop/peppermint/internal/grpc/client"
)

type HubClient interface {
	c.BaseClient
	hub.HubClient
	//JoinHelloTask(j job.Job) (job.Init, job.Run, job.Finalize)
	//JoinTask(j job.Job) (job.Init, job.Run, job.Finalize)
}

type hubClient struct {
	c.BaseClient
	hub.HubClient
}
