package webproxy

import (
	"github.com/AgentCoop/peppermint/cmd"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/service"
	g "github.com/AgentCoop/peppermint/internal/service/webproxy/grpc/server"
	model "github.com/AgentCoop/peppermint/internal/service/webproxy/model"
	service2 "github.com/AgentCoop/peppermint/pkg/service"
)

const (
	Name = "WebProxy"
)

type webProxy struct {
	service.WebProxyConfigurator
}

func init() {
	proxy := &webProxy{
		NewConfigurator(),
	}
	proxy.WebProxyConfigurator = NewConfigurator()
	reg := runtime.GlobalRegistry()
	serviceInfo := &runtime.ServiceInfo{
		Name: Name,
		Cfg: proxy.WebProxyConfigurator,
		Initializer: proxy.initializer,
	}
	reg.RegisterService(serviceInfo)
	reg.RegisterParserCmdHook(cmd.CMD_NAME_DB_MIGRATE, proxy.migrateDb)
}

func (w *webProxy) initializer() service2.Service {
	proxy := g.NewServer(
		Name,
		w.WebProxyConfigurator.Address(),
		w.WebProxyConfigurator,
		w,
	)
	return proxy
}

func (w *webProxy) migrateDb(options interface{}) {
	db := runtime.GlobalRegistry().Db()
	h := db.Handle()
	h.AutoMigrate(&model.WebProxyConfig{})
}

