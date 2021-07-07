package hub

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	api "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/service"
	grpc "github.com/AgentCoop/peppermint/internal/service/hub/grpc/server"
	"github.com/AgentCoop/peppermint/internal/service/hub/logger"
	"github.com/AgentCoop/peppermint/internal/service/hub/model"
	"github.com/AgentCoop/peppermint/pkg"
	grpc2 "github.com/AgentCoop/peppermint/pkg/grpc"
	svcPkg "github.com/AgentCoop/peppermint/pkg/service"
	"net"
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
	reg.RegisterHook(runtime.ServiceInitHook, func(args ...interface{}) error {
		return hub.Init()
	})
	reg.RegisterHook(runtime.CmdCreateDbHook, func(args ...interface{}) error {
		return hub.createDd(args...)
	})
}

func (hub *HubService) ReloadConfig(nodeId uint) error {
	cfg, err := hub.FetchConfig(nodeId)
	if err != nil {
		return err
	}
	hub.WithConfigurator(cfg)
	job.Logger(logger.Info)("%s configuration has been reloaded", hub.ShortName())
	return nil
}

func (hub *HubService) FetchConfig(nodeId uint) (svcPkg.ServiceConfigurator, error) {
	rt := runtime.GlobalRegistry().Runtime()
	app := runtime.GlobalRegistry().App().(pkg.AppNode)
	svcDb := app.ServiceDb()
	hubDb := model.NewDb(svcDb)
	cfg := model.NewConfigurator(hubDb)
	err := cfg.Fetch(nodeId)
	if err != nil {
		return nil, err
	}
	cfg.MergeCliOptions(rt.CliParser())
	return cfg, nil
}

func (hub *HubService) Init() error {
	rt := runtime.GlobalRegistry().Runtime()
	app := runtime.GlobalRegistry().App().(pkg.AppNode)
	var ipcSrv grpc2.BaseServer
	hub.Service = service.NewBaseService(Name)
//	hub.Service.OpenDb()

	// Service configurator
	cfg, err := hub.FetchConfig(app.Node().Id())
	if err != nil {
		return err
	}
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
	return nil
}

func (hub *HubService) encKeyStoreFallback(key interface{}) (interface{}, error) {
	nodeId := key.(i.NodeId)
	db := model.NewDb(hub.Db())
	node, err := db.FetchById(nodeId)
	if err != nil {
		return nil, err
	}
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

func (hub *HubService) createDd(args ...interface{}) error {
	force := args[0].(bool)
	svcDb := args[1].(pkg.Db)
	hub.Service = service.NewBaseService(Name)
	hubDb := model.NewDb(svcDb)
	if force {
		if err := hubDb.DropTables(); err != nil {
			return err
		}
	}
	return hubDb.CreateTables()
}
