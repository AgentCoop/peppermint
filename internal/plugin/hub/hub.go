
package hub

import (
	"github.com/AgentCoop/peppermint/cmd"
	i "github.com/AgentCoop/peppermint/internal"
	api "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	plugin "github.com/AgentCoop/peppermint/internal/plugin"
	grpc "github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/server"
	"github.com/AgentCoop/peppermint/internal/plugin/hub/model"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

const (
	Name = "Hub"
)

type hubService struct {
	server plugin.
	config.HubConfigurator
}

func init() {
	hub := new(hubService)
	hub.HubConfigurator = model.NewConfigurator()
	reg := runtime.GlobalRegistry()
	serviceInfo := &runtime.ServiceInfo{
		Name: Name,
		Cfg: hub.HubConfigurator,
		Initializer: hub.initializer,
	}
	reg.RegisterService(serviceInfo)
	reg.RegisterParserCmdHook(cmd.CMD_NAME_DB_MIGRATE, hub.migrateDb)
	reg.RegisterParserCmdHook(cmd.CMD_NAME_DB_CREATE, hub.createDd)
}

func (w *hubService) initializer() g.BaseServer {
	srv := grpc.NewServer(
		api.Hub_ServiceDesc.ServiceName,
		w.HubConfigurator.Address(),
	)
	w.server = srv
	w.RegisterEncKeyStoreFallback()
	return srv
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
	db := runtime.GlobalRegistry().Db()
	h := db.Handle()
	h.AutoMigrate(&model.HubConfig{}, &model.HubJoinedNode{}, &model.HubNodeTag{})
}

func (hub *hubService) createDd(options interface{}) {
	db := runtime.GlobalRegistry().Db()
	h := db.Handle()
	h.Migrator().DropTable(model.Tables...)
	h.Migrator().CreateTable(model.Tables...)
}

