package join

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/crypto"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	data "github.com/AgentCoop/peppermint/internal/grpc/data/hub/client/join"
)

func (ctx *joinCtx) JoinCmdTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		keyExch := crypto.NewKeyExchange(task)
		pubKey := keyExch.GetPublicKey()

		req := data.NewJoinHello(pubKey)
		ctx.joinHelloReqCh <- req
		resp := <-ctx.joinHelloResCh

		dataBag := resp.(data.JoinHello_DataBag)
		ctx.encKey = keyExch.ComputeKey(dataBag.HubPubKey())

		codec.SetEncKey(ctx.encKey)
		fmt.Printf("client enc key %x\n", ctx.encKey)
		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}


