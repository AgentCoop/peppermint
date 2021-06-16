package join

import (
	"context"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/crypto"
	"github.com/AgentCoop/peppermint/internal/runtime"

	//"github.com/AgentCoop/peppermint/internal/grpc/calldesc"
	"github.com/AgentCoop/peppermint/internal/model/node"
	cc "github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/client"
	//hh "github.com/AgentCoop/peppermint/internal/plugin/hub"
	//"github.com/AgentCoop/peppermint/internal/runtime"
)

func (c *joinContext) JoinTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		rt := runtime.GlobalRegistry().Runtime()
		ctx := context.Background()
		hubClient := j.GetValue().(cc.HubClient)

		// Generate a DH public key and exchange it with the hub server
		keyExch := crypto.NewKeyExchange(task)
		pubKey := keyExch.GetPublicKey()
		reqHello := &hub.JoinHello_Request{DhPubKey: pubKey}
		resHello, err := hubClient.JoinHello(ctx, reqHello)
		task.Assert(err)

		// Set computed encryption key for the client
		c.encKey = keyExch.ComputeKey(resHello.GetDhPubKey())
		err = node.UpdateNodeEncKey(c.encKey)
		task.Assert(err)
		rt.NodeConfigurator().Refresh()

		// Finish the join procedure
		reqJoin := &hub.Join_Request{
			Tag:           c.nodeTags,
			AvailServices: nil,
			Flags:         nil,
			JoinSecret:    c.secret,
		}
		ctx = context.Background()
		_, err = hubClient.Join(ctx, reqJoin)
		task.Assert(err)
		task.Done()
	}
	return nil, run, nil
}