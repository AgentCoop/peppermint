package join

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/crypto"
	"github.com/AgentCoop/peppermint/internal/grpc/session"
)

func (ctx *joinContext) JoinHelloTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		callDesc, ipc := session.Ipc_CallDesc(j, 0)

		// Extract data from the request
		req := callDesc.RequestData()
		dataBag := req.(joinHello_DataBag)

		// Compute encryption key
		pubKey := dataBag.NodePubKey()
		keyExch := crypto.NewKeyExchange(task)
		ctx.encKey = keyExch.ComputeKey(pubKey)

		resp := NewJoinHelloResponse(keyExch.GetPublicKey())
		callDesc.SetResponseData(resp)

		ipc.Svc_Send(0, nil)
		task.Done()
	}
	return nil, run, nil
}
