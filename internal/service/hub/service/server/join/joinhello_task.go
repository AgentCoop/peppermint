package join

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/crypto"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"github.com/AgentCoop/peppermint/internal/runtime"
	data "github.com/AgentCoop/peppermint/internal/service/hub/grpc/data/server/join"
)

func (ctx *joinCtx) JoinHelloTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		comm := j.GetValue().(runtime.GrpcServiceCommunicator)
		v := comm.ServiceRx(0)
		task.AssertNotNil(v)
		callDesc := v.(server.GrpcCallDescriptor)
		req := callDesc.GetRequest()

		dataBag := req.(data.DataBag)
		pubKey := dataBag.NodePubKey()

		keyExch := crypto.NewKeyExchange(task)
		ctx.encKey = keyExch.ComputeKey(pubKey)

		data.NewJoinHelloResponse(callDesc, keyExch.GetPublicKey())
		comm.ServiceTx(0, callDesc)
		task.Done()
	}
	return nil, run, nil
}
