package hub

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	c "github.com/AgentCoop/peppermint/internal/grpc/client"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/client"
	md_middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/client/metadata"
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

func (c *hubClient) withMiddlewares() grpc.DialOption {
	return middleware.WithUnaryClientChain(
		md_middleware.UnaryClientInterceptor(c),
	)
}

func NewClient(baseClient c.BaseClient) *hubClient {
	hubClient := new(hubClient)
	hubClient.BaseClient = baseClient
	hubClient.OnConnectedHook(func(cc grpc.ClientConnInterface) {
		hubClient.grpcHandle = hub.NewHubClient(cc)
	})
	hubClient.WithMiddlewareHook(func() grpc.DialOption {
		return hubClient.withMiddlewares()
	})
	return hubClient
}
