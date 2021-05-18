package join

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/crypto"
	data "github.com/AgentCoop/peppermint/internal/grpc/data/hub/join"
)

func (ctx *joinCtx) JoinHelloTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		pair := <-ctx.ReqChan[0]
		task.AssertNotNil(pair)
		req := pair.GetRequest()

		dataBag := req.(data.DataBag)
		pubKey := dataBag.NodePubKey()

		keyExch := crypto.NewKeyExchange(task)
		ctx.encKey = keyExch.ComputeKey(pubKey)

		data.NewJoinHelloResponse(pair, keyExch.GetPublicKey())
		ctx.ResChan[0] <- nil

		task.Done()
	}
	return nil, run, nil
}
