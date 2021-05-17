package hub

import (
	"context"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	jctx "github.com/AgentCoop/peppermint/internal/service/hub/client"
	data "github.com/AgentCoop/peppermint/internal/grpc/data/hub/client/join"
)

func (c *hubClient) JoinTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		joinCtx := j.GetValue().(jctx.JoinContext)
		req := <- joinCtx.JoinRequest()

		origResp, err := c.grpcHandle.Join(req.(context.Context), req.ToGrpcRequest().(*hub.Join_Request))
		task.Assert(err)

		resp := data.NewJoinResponse(req.(context.Context), origResp)
		joinCtx.JoinResponse() <- resp

		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

