package service

//import (
//	"github.com/AgentCoop/peppermint/internal/grpc"
//	"github.com/AgentCoop/peppermint/internal/grpc/service"
//	"github.com/AgentCoop/peppermint/internal/plugin/hub"
//	"github.com/AgentCoop/peppermint/internal/plugin/hub/model"
//	"github.com/AgentCoop/peppermint/internal/runtime"
//	"github.com/AgentCoop/peppermint/internal/runtime/deps"
//	"net"
//)
//
//func DefaultInit(svc grpc.Service, netSrv grpc.BaseServer, cfg deps.ServiceConfigurator) (grpc.Service, error) {
//	rt := runtime.GlobalRegistry().Runtime()
//	var ipcSrv grpc.BaseServer
//	policy := service.NewServicePolicy(netSrv.FullName(), netSrv.Methods())
//	if len(policy.Ipc_UnixDomainSocket()) > 0 {
//		unixAddr, _ := net.ResolveUnixAddr("unix", policy.Ipc_UnixDomainSocket())
//		ipcSrv = grpc.NewServer(svc.Name(), unixAddr)
//	}
//	// Configurator
//	cfg.Fetch()
//	cfg.MergeCliOptions(rt.CliParser())
//	hub.Service = service.NewBaseService(srv, ipcSrv, policy)
//	svc.RegisterEncKeyStoreFallback()
//	rt.RegisterService(Name, hub)
//	return hub, nil
//}
