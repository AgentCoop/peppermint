package service

import (
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/utils"
)

type baseService struct {
	srv    runtime.BaseServer
	ipcSrv runtime.BaseServer
	policy runtime.ServicePolicy
	cfg    runtime.ServiceConfigurator
}

func (b *baseService) RegisterEncKeyStoreFallback() {
	panic("implement me")
}

func (b *baseService) Name() string {
	return utils.Grpc_ExtractServerShortName(b.srv.FullName())
}

func (b *baseService) Server() runtime.BaseServer {
	return b.srv
}

func (b *baseService) IpcServer() runtime.BaseServer {
	return b.ipcSrv
}

func (b *baseService) Policy() runtime.ServicePolicy {
	return b.policy
}

func (b *baseService) Configurator() runtime.ServiceConfigurator {
	return b.cfg
}
