
package hub

import (
	"github.com/AgentCoop/peppermint/cmd"
	i "github.com/AgentCoop/peppermint/internal"
	grpc "github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/server"
	"github.com/AgentCoop/peppermint/internal/plugin/hub/model"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/config"
)

const (
	Name = "Hub"
)

type hubService struct {
	config.HubConfigurator
}

func init() {
	hub := &hubService{
		model.NewConfigurator(),
	}
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

func (w *hubService) initializer() runtime.Service {
	proxy := grpc.NewServer(
		Name,
		w.HubConfigurator.Address(),
	)
	w.RegisterEncKeyStoreFallback()
	return proxy
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

