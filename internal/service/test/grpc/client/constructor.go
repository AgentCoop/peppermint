package client

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/frontoffice/test"
	c "github.com/AgentCoop/peppermint/internal/grpc/client"
	"github.com/AgentCoop/peppermint/internal/runtime/service"
	tt "github.com/AgentCoop/peppermint/internal/service/test"
	"github.com/AgentCoop/peppermint/internal/utils"
	"google.golang.org/grpc"
)

func NewClient(target string) *testClient {
	testClient := new(testClient)
	testClient.BaseClient = c.NewBaseClient(target)
	connProvider := func(cc grpc.ClientConnInterface) {
		testClient.TestClient = test.NewTestClient(cc)
	}
	svcPolicy := service.NewServicePolicy(tt.Name, utils.Grpc_MethodsFromServiceDesc(test.Test_ServiceDesc))
	c.NewDefaultClient(testClient, svcPolicy, connProvider)
	return testClient
}
