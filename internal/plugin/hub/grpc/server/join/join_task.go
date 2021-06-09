package join

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc/session"
	"github.com/AgentCoop/peppermint/internal/plugin/hub/model"
)

func (ctx *joinContext) JoinTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		desc, ipc := session.Ipc_CallDesc(j, 1)
		req := desc.RequestData()
		nodeId := desc.Meta().NodeId()
		data := req.(*joinRequest)
		_ = data

		resp := NewJoinResponse()
		desc.SetResponseData(resp)

		err := model.AcceptJoin(nodeId)
		task.Assert(err)
		ipc.Svc_Send(1, nil)
		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

