package join

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/crypto"
	utils "github.com/AgentCoop/peppermint/internal/utils/grpc"
)

func (ctx *joinContext) JoinHelloTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		callDesc, ipc := utils.Ipc_CallDesc(j, 0)

		// Extract data from the request
		req := callDesc.GetRequest()
		dataBag := req.(joinHello_DataBag)

		// Compute encryption key
		pubKey := dataBag.NodePubKey()
		keyExch := crypto.NewKeyExchange(task)
		ctx.encKey = keyExch.ComputeKey(pubKey)

		NewJoinHelloResponse(callDesc, keyExch.GetPublicKey())
		ipc.Svc_Send(0, nil)
		task.Done()
	}
	return nil, run, nil
}
