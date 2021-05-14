package hub

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/crypto"
	"github.com/AgentCoop/peppermint/internal/grpc/client/hub"
)

type joinCtx struct {
	client hub.HubClient
	secret string
}

func (ctx *joinCtx) JoinTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		nodeKeyExchg := crypto.NewKeyExchange(task)
		hubPubKey := ctx.client.JoinHello(nodeKeyExchg.GetPublicKey())
		key := nodeKeyExchg.ComputeKey(hubPubKey)
		_ = key
	}
	return nil, run, nil
}