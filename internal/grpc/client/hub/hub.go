package hub

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	c "github.com/AgentCoop/peppermint/internal/grpc/client"
)

type HubClient interface {
	c.Client
	JoinHello(dfPublicKey []byte)
}

type hubClient struct {
	c.BaseClient
	grpcHandle hub.HubClient
}

func NewClient(address string, task job.Task) *hubClient {
	hubClient := &hubClient{}
	hubClient.Address = address
	hubClient.Task = task
	hubClient.Connect()
	hubClient.grpcHandle = hub.NewHubClient(hubClient.Conn)
	return hubClient
}
