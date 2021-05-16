package hub

import (
	"context"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	data "github.com/AgentCoop/peppermint/internal/grpc/data/hub/client/join"
	jctx "github.com/AgentCoop/peppermint/internal/service/hub/client"

	//"github.com/AgentCoop/peppermint/internal/service/hub/client/join"

	//"github.com/AgentCoop/peppermint/internal/api/peppermint/service"
)

func (c *hubClient) JoinHelloTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		ctx := context.Background()

		joinCtx := j.GetValue().(jctx.JoinContext)
		req := <- joinCtx.JoinHelloRequest()

		origResp, err := c.grpcHandle.JoinHello(ctx, req.ToGrpcRequest().(*hub.JoinHello_Request))
		task.Assert(err)

		resp := data.NewJoinHelloResponse(origResp)
		joinCtx.JoinHelloResponse() <- resp

		task.Done()
	}
	return nil, run, nil
}