package balancer

import (
	job "github.com/AgentCoop/go-work"
	api "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/balancer"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/service"
	g "github.com/AgentCoop/peppermint/internal/service/balancer/grpc/server"
	"github.com/AgentCoop/peppermint/internal/service/balancer/logger"
	"github.com/AgentCoop/peppermint/internal/service/balancer/model"
	service2 "github.com/AgentCoop/peppermint/pkg/service"
)

var (
	Name = api.Balancer_ServiceDesc.ServiceName
)

type lbService struct {
	service2.Service
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

func (lb *lbService) Init() (service2.Service, error) {
	rt := runtime.GlobalRegistry().Runtime()
	// Configurator
	cfg := model.NewConfigurator()
	cfg.Fetch()
	//cfg.MergeCliOptions(rt.CliParser())

	// Create network server and service policy
	srv := g.NewServer(Name, cfg.Address(), lb)
	srv.WithStdoutLogger(job.Logger(logger.Info))
	policy := service.NewServicePolicy(srv.FullName(), srv.Methods())

	lb.Service = service.NewBaseService(srv, nil, cfg, policy)
	rt.RegisterService(Name, lb)
	return lb, nil
}

//func (hub *lbService) encKeyStoreFallback(key interface{}) (interface{}, error) {
//	nodeId := key.(i.NodeId)
//	node, err := model.FetchById(nodeId);
//	if err != nil { return nil, err }
//	return node.encKey, nil
//}
//
//func (hub *lbService) RegisterEncKeyStoreFallback() {
//	rt := runtime.GlobalRegistry().Runtime()
//	rt.NodeManager().EncKeyStore().RegisterFallback(hub.encKeyStoreFallback)
//}

func (hub *lbService) migrateDb(options interface{}) {

}

func (hub *lbService) createDd(args...interface{}) {
	force := args[0].(bool)
	if force {
		model.DropTables()
	}
	model.CreateTables()
}
