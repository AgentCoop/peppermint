package client

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/frontoffice/test"
	c "github.com/AgentCoop/peppermint/internal/grpc/client"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/client"
	"github.com/AgentCoop/peppermint/internal/runtime"
	tt "github.com/AgentCoop/peppermint/internal/service/test"
	"google.golang.org/grpc"
	"net"
)

func NewCmdContext(opts interface{}, count int) *cmdContext {
	ctx := new(cmdContext)
	ctx.opts = opts
	if count == 0 {
		ctx.count = 1
	}
	return ctx
}

func NewJob(cmdName string, ctx *cmdContext, client *testClient) job.Job {
	j := job.NewJob(nil)
	j.AddOneshotTask(client.ConnectTask)
	switch cmdName {
	case CMD_NAME_PING:
		j.AddTask(ctx.PingTask)
	}
	return j
}

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
