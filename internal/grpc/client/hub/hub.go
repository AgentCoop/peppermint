package hub

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	c "github.com/AgentCoop/peppermint/internal/grpc/client"
	"google.golang.org/grpc"
)

type HubClient interface {
	c.BaseClient
	JoinHelloTask(j job.Job) (job.Init, job.Run, job.Finalize)
	JoinTask(j job.Job) (job.Init, job.Run, job.Finalize)
}

type hubClient struct {
	c.BaseClient
	grpcHandle hub.HubClient
}

func NewClient(baseClient c.BaseClient) *hubClient {
	hubClient := new(hubClient)
	hubClient.BaseClient = baseClient
	hubClient.OnConnectedHook(func(cc grpc.ClientConnInterface) {
		hubClient.grpcHandle = hub.NewHubClient(cc)
	})
	return hubClient
}
