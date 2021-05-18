package hub

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	data "github.com/AgentCoop/peppermint/internal/grpc/data/hub/client/join"
	jctx "github.com/AgentCoop/peppermint/internal/service/hub/client"
)

func (c *hubClient) JoinTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		joinCtx := j.GetValue().(jctx.JoinContext)
		pair := <- joinCtx.ReqChan(1)
		req := pair.GetRequest()

		origResp, err := c.grpcHandle.Join(pair, req.ToGrpcRequest().(*hub.Join_Request))
		task.Assert(err)

		data.NewJoinResponse(pair, origResp)
		joinCtx.ResChan(1) <- struct{}{}

		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

