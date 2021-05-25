package client

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	data "github.com/AgentCoop/peppermint/internal/service/hub/grpc/data/client/join"
	jctx "github.com/AgentCoop/peppermint/internal/service/hub/service/client"
	//"github.com/AgentCoop/peppermint/internal/service/hub/client/join"

	//"github.com/AgentCoop/peppermint/internal/api/peppermint/service"
)

func (c *hubClient) JoinHelloTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		joinCtx := j.GetValue().(jctx.JoinContext)
		pair := <- joinCtx.ReqChan(0)
		req := pair.GetRequest()

		origResp, err := c.grpcHandle.JoinHello(pair, req.ToGrpcRequest().(*hub.JoinHello_Request))
		task.Assert(err)

		data.NewJoinHelloResponse(pair, origResp)
		joinCtx.ResChan(0) <- struct{}{}

		task.Done()
	}
	return nil, run, nil
}