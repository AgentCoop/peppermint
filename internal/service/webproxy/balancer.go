package webproxy

import (
	client2 "github.com/AgentCoop/peppermint/internal/grpc/client"
	"github.com/AgentCoop/peppermint/internal/grpc/proxy"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc"
)

func (b *webProxy) SimpleRandom(svcName string, pool runtime.NodePool) runtime.ServiceEndpoint {
	return nil
}

func (b *webProxy) ForwardCall(srv interface{}, stream grpc.ServerStream) error {
	fullMethod, _ := grpc.MethodFromServerStream(stream)
	locator := runtime.GlobalRegistry().ServiceLocator("")
	nodesPool := locator.FindByMethodName(fullMethod)
	available := nodesPool.FilterByStatus(runtime.Available)
	if available.Len() == 0 {
		//return status.Status(codes.Unavailable, "12")
	}
	proxiedSvcName := locator.ServiceNameByMethod(fullMethod)
	endpoint := b.SimpleRandom(proxiedSvcName, available)
	c := client2.NewBaseClient(endpoint)
	pjob, _ := proxy.NewProxyConnJob(c, stream, nil)
	<-pjob.Run()
	return nil
}
