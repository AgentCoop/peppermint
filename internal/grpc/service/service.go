package service

import (
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/utils"
)

type baseService struct {
	srv    grpc.BaseServer
	policy grpc.ServicePolicy
}

func (b *baseService) Name() string {
	return utils.Grpc_ExtractServerShortName(b.srv.FullName())
}

func (b *baseService) Server() grpc.BaseServer {
	return b.srv
}

func (b *baseService) Policy() grpc.ServicePolicy {
	return b.policy
}
