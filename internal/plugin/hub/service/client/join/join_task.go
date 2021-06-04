package join

import (
	"context"
	"fmt"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/crypto"
	"github.com/AgentCoop/peppermint/internal/grpc/client"
	data "github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/data/client/join"
)

func (ctx *joinCtx) JoinCmdTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
	}
	run := func(task job.Task) {
		pair := client.NewRequestResponsePair(ctx.HubClient, context.Background())

		keyExch := crypto.NewKeyExchange(task)
		pubKey := keyExch.GetPublicKey()

		data.NewJoinHello(pair, pubKey)
		ctx.ReqChan(0) <- pair
		<-ctx.ResChan(0)

		resp := pair.GetResponse()
		dataBag := resp.(data.JoinHello_DataBag)
		ctx.encKey = keyExch.ComputeKey(dataBag.HubPubKey())

		//codec.SetEncKey(ctx.encKey)
		fmt.Printf("client enc key %x\n", ctx.encKey)

		pair = client.NewRequestResponsePair(ctx.HubClient, context.Background())
		data.NewJoin(pair, ctx.secret)
		ctx.ReqChan(1) <- pair

		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}


