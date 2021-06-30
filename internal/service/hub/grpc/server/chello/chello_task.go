package chello

import (
	job "github.com/AgentCoop/go-work"
	s "github.com/AgentCoop/peppermint/internal/security"
	"github.com/AgentCoop/peppermint/internal/grpc/session"
	"github.com/AgentCoop/peppermint/internal/utils"
)

func (ctx *clientHelloCtx) ClientHelloTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		desc, ipc := session.Ipc_CallDesc(j, task.Index()-1)

		// Compute encryption key and return hub's public key
		req := desc.RequestData()
		dataBag := req.(clientHelloData)
		pubKey := dataBag.NodePubKey()

		keyExch, err := s.NewKeyExchange()
		task.Assert(err)

		ctx.randMsg = utils.Rand_BytesArray(32)
		ctx.encKey, err = keyExch.ComputeKey(pubKey)
		task.Assert(err)

		resp := NewClientHelloResponse(keyExch.GetPublicKey(), ctx.randMsg)
		desc.SetResponseData(resp)

		ipc.Svc_Send(task.Index()-1, nil)
		task.Done()
	}
	return nil, run, nil
}
