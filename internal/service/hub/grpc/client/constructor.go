package client

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	c "github.com/AgentCoop/peppermint/internal/grpc/client"
	"github.com/AgentCoop/peppermint/internal/runtime/service"
	hh "github.com/AgentCoop/peppermint/internal/service/hub"
	"github.com/AgentCoop/peppermint/internal/utils"
	"google.golang.org/grpc"
)

func NewClient(target string) *hubClient {
	hubClient := new(hubClient)
	hubClient.BaseClient = c.NewBaseClient(target)
	connProvider := func(cc grpc.ClientConnInterface) {
		hubClient.HubClient = hub.NewHubClient(cc)
	}
	svcPolicy := service.NewServicePolicy(hh.Name, utils.Grpc_MethodsFromServiceDesc(hub.Hub_ServiceDesc))
	c.NewDefaultClient(hubClient, svcPolicy, connProvider)
	return hubClient
}
