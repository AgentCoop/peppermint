package join

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/crypto"
	data "github.com/AgentCoop/peppermint/internal/grpc/data/hub/join"
	srv "github.com/AgentCoop/peppermint/internal/grpc/server"
)

func (ctx *joinCtx) JoinHelloTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		req := <-ctx.JoinHelloReqCh
		task.AssertNotNil(req)

		dataBag := req.(data.DataBag)
		pubKey := dataBag.NodePubKey()

		keyExch := crypto.NewKeyExchange(task)
		ctx.encKey = keyExch.ComputeKey(pubKey)

		res := data.NewJoinHelloResponse(req.(srv.MetaData), keyExch.GetPublicKey())
		ctx.JoinHelloRespCh <- res

		task.Done()
	}
	return nil, run, nil
}
