package join

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/crypto"
	data "github.com/AgentCoop/peppermint/internal/grpc/data/hub/join"
)

func (ctx *joinCtx) JoinHelloTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		req := <-ctx.JoinHelloReqCh
		task.AssertNotNil(req)

		dataBag := req.(data.DataBag)
		pubKey := dataBag.NodePubKey()

		nodeKeyExchng := crypto.NewKeyExchange(task)
		ctx.encKey = nodeKeyExchng.ComputeKey(pubKey)

		res := data.NewJoinHelloResponse(nil)
		ctx.JoinHelloRespCh <- res

		task.Done()
	}
	return nil, run, nil
}
