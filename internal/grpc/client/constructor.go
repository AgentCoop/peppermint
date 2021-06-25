package client

import (
	g "github.com/AgentCoop/peppermint/internal/grpc"
	middleware "github.com/AgentCoop/peppermint/internal/grpc/middleware/client"
	"github.com/AgentCoop/peppermint/pkg/service"
	"google.golang.org/grpc"
)

func NewBaseClient(target string, opts ...grpc.DialOption) *baseClient {
	c := new(baseClient)
	c.target = target
	c.opts = opts
	return c
}

func NewDefaultClient(cc g.BaseClient, svcPolicy service.ServicePolicy, connProvider g.ConnProvider) {
	cc.WithConnProvider(connProvider)
	if svcPolicy.DefaultPort() > 0 {
		cc.WithTargetPort(svcPolicy.DefaultPort())
	}
	cc.WithUnaryInterceptors(
		middleware.PreUnaryInterceptor(cc, svcPolicy),
		middleware.SecureChannelUnaryInterceptor(),
		middleware.PostUnaryInterceptor(cc, svcPolicy),
	)
}
