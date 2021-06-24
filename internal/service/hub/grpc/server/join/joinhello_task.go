package join

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/crypto"
	"github.com/AgentCoop/peppermint/internal/grpc/session"
	"github.com/AgentCoop/peppermint/internal/service/hub/model"
)

func (ctx *joinContext) JoinHelloTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		desc, ipc := session.Ipc_CallDesc(j, 0)
		// Joined nodes have to call the disjoin method in order to join again
		nodeId := desc.Meta().NodeId()
		if model.HasJoined(nodeId) {
			j.Cancel(errAlreadyJoined)
			return
		}
		// Compute encryption key and return hub's public key
		req := desc.RequestData()
		dataBag := req.(joinHello_DataBag)
		pubKey := dataBag.NodePubKey()

		keyExch, err := crypto.NewKeyExchange()
		task.Assert(err)

		ctx.encKey, err = keyExch.ComputeKey(pubKey)
		task.Assert(err)

		resp := NewJoinHelloResponse(keyExch.GetPublicKey())
		desc.SetResponseData(resp)
		err = model.SaveJoinRequest(nodeId, ctx.encKey)
		task.Assert(err)

		ipc.Svc_Send(0, nil)
		task.Done()
	}
	return nil, run, nil
}
