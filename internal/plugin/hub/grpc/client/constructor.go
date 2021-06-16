package client

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	c "github.com/AgentCoop/peppermint/internal/grpc/client"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/client"
	hh "github.com/AgentCoop/peppermint/internal/plugin/hub"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc"
	"net"
)

func NewClient(addr net.Addr, opts ...grpc.DialOption) *hubClient {
	hubClient := new(hubClient)
	hubClient.BaseClient = c.NewBaseClient(addr, opts...)
	hubClient.WithConnProvider(func(cc grpc.ClientConnInterface) {
		hubClient.HubClient = hub.NewHubClient(cc)
	})
	rt := runtime.GlobalRegistry().Runtime()
	svcPolicy :=rt.ServicePolicyByName(hh.Name)
	hubClient.WithUnaryInterceptors(
		middleware.PreUnaryInterceptor(svcPolicy),
		middleware.SecureChannelUnaryInterceptor(),
		middleware.PostUnaryInterceptor(hubClient),
	)
	return hubClient
}
