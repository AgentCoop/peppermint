package client

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	c "github.com/AgentCoop/peppermint/internal/grpc/client"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/client"
	"google.golang.org/grpc"
	"net"
)

func NewClient(addr net.Addr, opts ...grpc.DialOption) *hubClient {
	hubClient := new(hubClient)
	hubClient.BaseClient = c.NewBaseClient(addr, opts...)
	hubClient.WithConnProvider(func(cc grpc.ClientConnInterface) {
		hubClient.HubClient = hub.NewHubClient(cc)
	})
	hubClient.WithUnaryInterceptors(
		middleware.PreUnaryInterceptor(hubClient),
		middleware.SecureChannelUnaryInterceptor(codec.Serialized),
		middleware.PostUnaryInterceptor(hubClient),
	)
	return hubClient
}
