package client

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/frontoffice/test"
	c "github.com/AgentCoop/peppermint/internal/grpc/client"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/client"
	"github.com/AgentCoop/peppermint/internal/runtime"
	tt "github.com/AgentCoop/peppermint/internal/service/test"
	"google.golang.org/grpc"
	"net"
)

func NewClient(addr net.Addr, opts ...grpc.DialOption) *testClient {
	testClient := new(testClient)
	testClient.BaseClient = c.NewBaseClient(addr, opts...)
	testClient.WithConnProvider(func(cc grpc.ClientConnInterface) {
		testClient.TestClient = test.NewTestClient(cc)
	})
	rt := runtime.GlobalRegistry().Runtime()
	svcPolicy := rt.ServicePolicyByName(tt.Name)
	testClient.WithUnaryInterceptors(
		middleware.PreUnaryInterceptor(testClient, svcPolicy),
		middleware.SecureChannelUnaryInterceptor(),
		middleware.PostUnaryInterceptor(testClient, svcPolicy),
	)
	return testClient
}
