package balancer

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	api "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/balancer"
	g "github.com/AgentCoop/peppermint/internal/plugin/balancer/grpc/server"
	"github.com/AgentCoop/peppermint/internal/plugin/balancer/logger"
	"github.com/AgentCoop/peppermint/internal/plugin/hub/model"
	//"github.com/AgentCoop/peppermint/internal/plugin/balancer/model"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/service"
)

var (
	Name = api.Balancer_ServiceDesc.ServiceName
)

type lbService struct {
	runtime.Service
}

func init() {
	lb := new(lbService)
	reg := runtime.GlobalRegistry()
	reg.RegisterHook(runtime.ServiceInitHook, func(args...interface{}) {
		lb.Init()
	})
	reg.RegisterHook(runtime.CmdCreateDbHook, func(args...interface{}) {
		lb.createDd(args...)
	})
}

func (lb *lbService) Init() (runtime.Service, error) {
	rt := runtime.GlobalRegistry().Runtime()
	var ipcSrv runtime.BaseServer
	// Configurator
	cfg := model.NewConfigurator()
	//cfg.Fetch()
	//cfg.MergeCliOptions(rt.CliParser())

	// Create network server and service policy
	srv := g.NewServer(Name, cfg.Address(), lb)
	srv.WithStdoutLogger(job.Logger(logger.Info))
	policy := service.NewServicePolicy(srv.FullName(), srv.Methods())

	// IPC server
	//if len(policy.Ipc_UnixDomainSocket()) > 0 {
	//	unixAddr, _ := net.ResolveUnixAddr("unix", policy.Ipc_UnixDomainSocket())
		//ipcSrv = grpc.NewServer(Name, unixAddr)
		//ipcSrv.WithStdoutLogger(job.Logger(logger.Info))
	//}
	lb.Service = service.NewBaseService(srv, ipcSrv, cfg, policy)
	lb.RegisterEncKeyStoreFallback()
	rt.RegisterService(Name, lb)
	return lb, nil
}

func (hub *lbService) encKeyStoreFallback(key interface{}) (interface{}, error) {
	nodeId := key.(i.NodeId)
	node, err := model.FetchById(nodeId);
	if err != nil { return nil, err }
	return node.EncKey, nil
}

func (hub *lbService) RegisterEncKeyStoreFallback() {
	rt := runtime.GlobalRegistry().Runtime()
	rt.NodeManager().EncKeyStore().RegisterFallback(hub.encKeyStoreFallback)
}

func (hub *lbService) migrateDb(options interface{}) {

}

func (hub *lbService) createDd(args...interface{}) {
	force := args[0].(bool)
	if force {
		model.DropTables()
	}
	model.CreateTables()
}
