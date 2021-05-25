
package hub

import (
	"github.com/AgentCoop/peppermint/cmd"
	model "github.com/AgentCoop/peppermint/internal/model/hub"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/config"
	"github.com/AgentCoop/peppermint/internal/service"
	grpc "github.com/AgentCoop/peppermint/internal/service/hub/grpc/server"
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
}

func (w *hubService) initializer() service.Service {
	proxy := grpc.NewServer(
		Name,
		w.HubConfigurator.Address(),
	)
	return proxy
}

func (hub *hubService) migrateDb(options interface{}) {
	db := runtime.GlobalRegistry().Db()
	h := db.Handle()
	h.AutoMigrate(&model.HubConfig{})
}


