package service

import (
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/db"
	"github.com/AgentCoop/peppermint/internal/utils"
	"github.com/AgentCoop/peppermint/pkg"
	"github.com/AgentCoop/peppermint/pkg/grpc"
	"github.com/AgentCoop/peppermint/pkg/service"
	"path"
	"strings"
)

type baseService struct {
	node   pkg.Node
	name   string
	srv    grpc.BaseServer
	ipcSrv grpc.BaseServer
	policy service.ServicePolicy
	cfg    service.ServiceConfigurator
	db     pkg.Db
}

func (b *baseService) OpenDb() error {
	// <nodeId>_<service_name>.db
	app := runtime.GlobalRegistry().App().(pkg.AppNode)
	basename := utils.Conv_IntToHex(b.node.ExternalId(), internal.NodeId(0).Size())
	basename += "_" + strings.ToLower(b.name) + ".db"
	pathname := path.Join(app.RootDir(), "db", basename)
	db, err := db.Open(pathname)
	b.db = db
	return err
}

//func (b *baseService) App() pkg.AppNode {
//	return b.app
//}

func (b *baseService) Db() pkg.Db {
	return b.db
}

func (b *baseService) RegisterEncKeyStoreFallback() {
	panic("implement me")
}

func (b *baseService) Name() string {
	return utils.Grpc_ExtractServerShortName(b.srv.FullName())
}

func (b *baseService) Server() grpc.BaseServer {
	return b.srv
}

func (b *baseService) WithServer(srv grpc.BaseServer) {
	b.srv = srv
}

func (b *baseService) IpcServer() grpc.BaseServer {
	return b.ipcSrv
}

func (b *baseService) WithIpcServer(ipcSrv grpc.BaseServer) {
	b.ipcSrv = ipcSrv
}

func (b *baseService) Policy() service.ServicePolicy {
	return b.policy
}

func (b *baseService) WithPolicy(policy service.ServicePolicy) {
	b.policy = policy
}

func (b *baseService) Configurator() service.ServiceConfigurator {
	return b.cfg
}

func (b *baseService) WithConfigurator(cfg service.ServiceConfigurator) {
	b.cfg = cfg
}
