
package hub

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	api "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/runtime/service"
	"github.com/AgentCoop/peppermint/internal/service/hub/logger"
	"github.com/AgentCoop/peppermint/pkg"
	grpc2 "github.com/AgentCoop/peppermint/pkg/grpc"
	svcPkg "github.com/AgentCoop/peppermint/pkg/service"
	"net"
	grpc "github.com/AgentCoop/peppermint/internal/service/hub/grpc/server"
	"github.com/AgentCoop/peppermint/internal/service/hub/model"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

var (
	Name = api.Hub_ServiceDesc.ServiceName
)

type HubService struct {
	svcPkg.Service
}

func init() {
	hub := new(HubService)
	reg := runtime.GlobalRegistry()
	reg.RegisterHook(runtime.ServiceInitHook, func(args...interface{}) {
		hub.Init()
	})
	reg.RegisterHook(runtime.CmdCreateDbHook, func(args...interface{}) {
		hub.createDd(args...)
	})
}

func (hub *HubService) Init() (svcPkg.Service, error) {
	rt := runtime.GlobalRegistry().Runtime()
	app := runtime.GlobalRegistry().App().(pkg.AppNode)
	var ipcSrv grpc2.BaseServer
	hub.Service = service.NewBaseService(Name, app.Node())
	hub.Service.OpenDb()

	// Service configurator
	hubDb := model.NewDb(hub.Service.Db())
	cfg := model.NewConfigurator(hubDb)
	cfg.Fetch(app.Node().Id())
	cfg.MergeCliOptions(rt.CliParser())
	hub.Service.WithConfigurator(cfg)

	// Create network server
	srv := grpc.NewServer(Name, cfg.Address())
	srv.WithStdoutLogger(job.Logger(logger.Info))
	hub.Service.WithServer(srv)

	// Set up service policy
	policy := service.NewServicePolicy(srv.FullName(), srv.Methods())
	hub.Service.WithPolicy(policy)

	// Create IPC server
	if len(policy.Ipc_UnixDomainSocket()) > 0 {
		unixAddr, _ := net.ResolveUnixAddr("unix", policy.Ipc_UnixDomainSocket())
		ipcSrv = grpc.NewServer(Name, unixAddr)
		ipcSrv.WithStdoutLogger(job.Logger(logger.Info))
	}
	hub.RegisterEncKeyStoreFallback()
	rt.RegisterService(Name, hub)
	return hub, nil
}

func (hub *HubService) encKeyStoreFallback(key interface{}) (interface{}, error) {
	nodeId := key.(i.NodeId)
	db := model.NewDb(hub.Db())
	node, err := db.FetchById(nodeId);
	if err != nil { return nil, err }
	return node.EncKey, nil
}

func (hub *HubService) RegisterEncKeyStoreFallback() {
	rt := runtime.GlobalRegistry().Runtime()
	rt.EncKeyStore().RegisterFallback(hub.encKeyStoreFallback)
}

func (hub *HubService) migrateDb(options interface{}) {
	//db := runtime.GlobalRegistry().Db()
	//h := db.Handle()
	//h.AutoMigrate(&model.HubConfig{}, &model.HubNode{}, &model.HubNodeTag{})
}

func (hub *HubService) createDd(args...interface{}) {
	force := args[0].(bool)
	node := args[1].(pkg.Node)
	hub.Service = service.NewBaseService(Name, node)
	err := hub.Service.OpenDb()
	if err != nil {
		// @todo rewrite
		panic("failed to open db")
	}
	hubDb := model.NewDb(hub.Db())
	if force {
		hubDb.DropTables()
	}
	hubDb.CreateTables()
}
