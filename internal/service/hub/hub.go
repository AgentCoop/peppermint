
package hub

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	api "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/service/hub/logger"
	"github.com/AgentCoop/peppermint/internal/runtime/service"
	"net"

	//plugin "github.com/AgentCoop/peppermint/internal/plugin"
	grpc "github.com/AgentCoop/peppermint/internal/service/hub/grpc/server"
	"github.com/AgentCoop/peppermint/internal/service/hub/model"
	//"github.com/AgentCoop/peppermint/internal/plugin/webproxy/model"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

var (
	Name = api.Hub_ServiceDesc.ServiceName
)

type hubService struct {
	runtime.Service
}

func init() {
	hub := new(hubService)
	reg := runtime.GlobalRegistry()
	reg.RegisterHook(runtime.ServiceInitHook, func(args...interface{}) {
		hub.Init()
	})
	reg.RegisterHook(runtime.CmdCreateDbHook, func(args...interface{}) {
		hub.createDd(args...)
	})
}

func (hub *hubService) Init() (runtime.Service, error) {
	rt := runtime.GlobalRegistry().Runtime()
	var ipcSrv runtime.BaseServer
	// Configurator
	cfg := model.NewConfigurator()
	cfg.Fetch()
	cfg.MergeCliOptions(rt.CliParser())

	// Create network server and service policy
	srv := grpc.NewServer(Name, cfg.Address())
	srv.WithStdoutLogger(job.Logger(logger.Info))
	policy := service.NewServicePolicy(srv.FullName(), srv.Methods())

	// IPC server
	if len(policy.Ipc_UnixDomainSocket()) > 0 {
		unixAddr, _ := net.ResolveUnixAddr("unix", policy.Ipc_UnixDomainSocket())
		ipcSrv = grpc.NewServer(Name, unixAddr)
		ipcSrv.WithStdoutLogger(job.Logger(logger.Info))
	}
	hub.Service = service.NewBaseService(srv, ipcSrv, cfg, policy)
	hub.RegisterEncKeyStoreFallback()
	rt.RegisterService(Name, hub)
	return hub, nil
}

func (hub *hubService) encKeyStoreFallback(key interface{}) (interface{}, error) {
	nodeId := key.(i.NodeId)
	node, err := model.FetchById(nodeId);
	if err != nil { return nil, err }
	return node.EncKey, nil
}

func (hub *hubService) RegisterEncKeyStoreFallback() {
	rt := runtime.GlobalRegistry().Runtime()
	rt.NodeManager().EncKeyStore().RegisterFallback(hub.encKeyStoreFallback)
}

func (hub *hubService) migrateDb(options interface{}) {
	//db := runtime.GlobalRegistry().Db()
	//h := db.Handle()
	//h.AutoMigrate(&model.HubConfig{}, &model.HubJoinedNode{}, &model.HubNodeTag{})
}

func (hub *hubService) createDd(args...interface{}) {
	force := args[0].(bool)
	if force {
		model.DropTables()
	}
	model.CreateTables()
}

