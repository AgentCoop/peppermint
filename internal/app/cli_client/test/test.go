package test

import (
	"context"
	t "github.com/AgentCoop/peppermint/internal/api/peppermint/service/frontoffice/test"
	"github.com/AgentCoop/peppermint/pkg"
	"google.golang.org/grpc"
)

type callParams struct {
	token          string
	repeat         uint
	workers        uint
	rsDelay        uint32
	rsDelayJitter  uint32
	rqBulkMin      uint32
	rqBulkMax      uint32
	callFailProbab uint
}

type ClientCallFn func(context.Context, *t.Request, ...grpc.CallOption) (*t.Response, error)

type appTest struct {
	pkg.App
	callParams  *callParams
	execCmdChan chan ClientCallFn
}
