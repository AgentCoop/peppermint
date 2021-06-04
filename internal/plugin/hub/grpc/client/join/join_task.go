package join

import (
	"context"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/crypto"
	"github.com/AgentCoop/peppermint/internal/model/node"
	cc "github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/client"
)

func (c *joinContext) JoinTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		ctx := context.Background()
		hubClient := j.GetValue().(cc.HubClient)

		// Generate DH public key
		keyExch := crypto.NewKeyExchange(task)
		pubKey := keyExch.GetPublicKey()

		reqHello := &hub.JoinHello_Request{
			DhPubKey: pubKey,
		}
		resHello, err := hubClient.JoinHello(ctx, reqHello)
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
		resJoin, err := hubClient.Join(ctx, reqJoin)
		task.Assert(err)
		_ = resJoin

		// Persist data of the newly joined node
		node.UpdateNode(c.encKey)
		task.Done()
	}
	return nil, run, nil
}