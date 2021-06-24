package service

import (
	"github.com/AgentCoop/peppermint/internal/utils"
	"github.com/AgentCoop/peppermint/pkg/grpc"
	"github.com/AgentCoop/peppermint/pkg/service"
)

type baseService struct {
	srv    grpc.BaseServer
	ipcSrv grpc.BaseServer
	policy service.ServicePolicy
	cfg    service.ServiceConfigurator
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

func (b *baseService) IpcServer() grpc.BaseServer {
	return b.ipcSrv
}

func (b *baseService) Policy() service.ServicePolicy {
	return b.policy
}

func (b *baseService) Configurator() service.ServiceConfigurator {
	return b.cfg
}
