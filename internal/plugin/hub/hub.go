
package hub

import (
	"github.com/AgentCoop/peppermint/cmd"
	i "github.com/AgentCoop/peppermint/internal"
	api "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/service"
	"net"

	//plugin "github.com/AgentCoop/peppermint/internal/plugin"
	grpc "github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/server"
	"github.com/AgentCoop/peppermint/internal/plugin/hub/model"
	//"github.com/AgentCoop/peppermint/internal/plugin/webproxy/model"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

var (
	Name = api.Hub_ServiceDesc.ServiceName
)

type hubService struct {
	g.Service
	model.HubConfigurator
}

func init() {
	hub := new(hubService)
	reg := runtime.GlobalRegistry()
	reg.RegisterHook(runtime.OnServiceInitHook, func(args...interface{}) {
		hub.Init()
	})
	reg.RegisterParserCmdHook(cmd.CMD_NAME_DB_MIGRATE, hub.migrateDb)
	reg.RegisterParserCmdHook(cmd.CMD_NAME_DB_CREATE, hub.createDd)
}

func (hub *hubService) Init() (g.Service, error) {
	rt := runtime.GlobalRegistry().Runtime()
	var ipcSrv g.BaseServer
	srv := grpc.NewServer(
		Name,
		hub.HubConfigurator.Address(),
	)
	policy := service.NewServicePolicy(srv.FullName(), srv.Methods())
	if len(policy.Ipc_UnixDomainSocket()) > 0 {
		unixAddr, _ := net.ResolveUnixAddr("unix", policy.Ipc_UnixDomainSocket())
		ipcSrv = grpc.NewServer(Name, unixAddr)
	}
	// Configurator
	hub.HubConfigurator = model.NewConfigurator()
	hub.HubConfigurator.Fetch()
	hub.HubConfigurator.MergeCliOptions(rt.CliParser())
	hub.Service = service.NewBaseService(srv, ipcSrv, policy)
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

func (hub *hubService) createDd(options interface{}) {
	//db := runtime.GlobalRegistry().Db()
	//h := db.Handle()
	//h.Migrator().DropTable(model.Tables...)
	//h.Migrator().CreateTable(model.Tables...)
}

