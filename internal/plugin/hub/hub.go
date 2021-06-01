
package hub

import (
	"github.com/AgentCoop/peppermint/cmd"
	"github.com/AgentCoop/peppermint/internal/plugin/hub/model"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/config"
	grpc "github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/server"
)

const (
	Name = "Hub"
)

type hubService struct {
	config.HubConfigurator
}

func init() {
	hub := &hubService{
		NewConfigurator(),
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
	return proxy
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

