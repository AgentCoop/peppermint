package client

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	c "github.com/AgentCoop/peppermint/internal/grpc/client"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/client"
	"google.golang.org/grpc"
)


func NewClient(baseClient c.BaseClient) *hubClient {
	hubClient := new(hubClient)
	hubClient.BaseClient = baseClient
	hubClient.WithConnProvider(func(cc grpc.ClientConnInterface) {
		hubClient.HubClient = hub.NewHubClient(cc)
	})
	hubClient.WithUnaryInterceptors(
		middleware.PreUnaryInterceptor(hubClient),
		middleware.SecureChannelUnaryInterceptor(hubClient),
		middleware.PostUnaryInterceptor(hubClient),
	)
	return hubClient
}
