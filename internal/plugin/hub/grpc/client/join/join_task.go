package join

import (
	"context"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/crypto"
	"github.com/AgentCoop/peppermint/internal/grpc/calldesc"
	"github.com/AgentCoop/peppermint/internal/model/node"
	cc "github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/client"
)

func (c *joinContext) JoinTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		ctx := context.Background()
		hubClient := j.GetValue().(cc.HubClient)

		// Generate a DH public key and exchange it with the hub server
		keyExch := crypto.NewKeyExchange(task)
		pubKey := keyExch.GetPublicKey()

		// Encryption key will be received after the call, make it insecure
		reqHello := &hub.JoinHello_Request{
			DhPubKey: pubKey,
		}
		callHelloDesc := calldesc.NewClientInSecure(ctx)
		resHello, err := hubClient.JoinHello(callHelloDesc, reqHello)
		task.Assert(err)

		// Set computed encryption key for the client
		c.encKey = keyExch.ComputeKey(resHello.GetDhPubKey())
		hubClient.WithEncKey(c.encKey)

		// Finish the join procedure
		reqJoin := &hub.Join_Request{
			Tag:           c.nodeTags,
			AvailServices: nil,
			Flags:         nil,
			JoinSecret:    c.secret,
		}
		ctx = context.Background()
		secPolicy := calldesc.NewSecurityPolicy(true, c.encKey)
		callDesc := calldesc.NewClient(ctx, secPolicy)
		callDesc.WithSessionFrom(callHelloDesc)
		resJoin, err := hubClient.Join(callDesc, reqJoin)
		task.Assert(err)
		_ = resJoin

		// Persist data of the newly joined node
		node.UpdateNode(c.encKey)
		task.Done()
	}
	return nil, run, nil
}