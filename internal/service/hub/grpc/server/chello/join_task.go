package chello

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc/session"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/service/hub/logger"
	"github.com/AgentCoop/peppermint/internal/service/hub/model"
	"github.com/AgentCoop/peppermint/pkg"
)

func (ctx *clientHelloCtx) JoinTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		app := runtime.GlobalRegistry().App().(pkg.AppNode)
		desc, ipc := session.Ipc_CallDesc(j, task.Index()-1)
		req := desc.RequestData().(*joinRequest)
		nodeId := desc.Meta().NodeId()

		svc := desc.Service()
		cfg := svc.Configurator().(model.HubConfigurator)
		invalidCreds := cfg.Secret() != req.secret
		task.AssertTrue(invalidCreds, errInvalidCreds)

		db := model.NewDb(app.ServiceDb())
		err := db.SaveJoinRequest(nodeId, ctx.encKey)
		task.Assert(err)

		job.Logger(logger.Info)("node #%x join accepted", nodeId)
		resp := NewJoinResponse()
		desc.SetResponseData(resp)

		ipc.Svc_Send(task.Index()-1, nil)
		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

