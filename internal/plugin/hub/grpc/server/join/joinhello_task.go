package join

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/crypto"
	"github.com/AgentCoop/peppermint/internal/grpc/session"
	"github.com/AgentCoop/peppermint/internal/plugin/hub/model"
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

		err := model.SaveJoinRequest(callDesc.Meta().NodeId(), ctx.encKey)
		task.Assert(err)
		ipc.Svc_Send(0, nil)
		task.Done()
	}
	return nil, run, nil
}
