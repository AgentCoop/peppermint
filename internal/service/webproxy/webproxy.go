package webproxy

import (
	"github.com/AgentCoop/peppermint/cmd"
	grpc "github.com/AgentCoop/peppermint/internal/grpc/webproxy"
	model "github.com/AgentCoop/peppermint/internal/model/webproxy"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/config"
	"github.com/AgentCoop/peppermint/internal/service"
)

const (
	Name = "WebProxy"
)

type webProxy struct {
	config.WebProxyConfigurator
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

func (w *webProxy) initializer() service.Service {
	proxy := grpc.NewServer(
		Name,
		w.WebProxyConfigurator.Address(),
		w.WebProxyConfigurator.ServerName(),
		w.WebProxyConfigurator.X509CertPEM(),
		w.WebProxyConfigurator.X509KeyPEM(),
	)
	return proxy
}

func (w *webProxy) migrateDb(options interface{}) {
	db := runtime.GlobalRegistry().Db()
	h := db.Handle()
	h.AutoMigrate(&model.WebProxyConfig{})
}

