package join

import (
	"context"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/crypto"
	//"github.com/AgentCoop/peppermint/internal/grpc/calldesc"
	"github.com/AgentCoop/peppermint/internal/model/node"
	cc "github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/client"
	//hh "github.com/AgentCoop/peppermint/internal/plugin/hub"
	//"github.com/AgentCoop/peppermint/internal/runtime"
)

func (c *joinContext) JoinTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		//rt := runtime.GlobalRegistry().Runtime()
		//svcPolicy := rt.ServicePolicyByName(hh.Name)
		ctx := context.Background()
		hubClient := j.GetValue().(cc.HubClient)

		// Generate a DH public key and exchange it with the hub server
		keyExch := crypto.NewKeyExchange(task)
		pubKey := keyExch.GetPublicKey()

		// Encryption key will be received after the call, make it insecure
		reqHello := &hub.JoinHello_Request{
			DhPubKey: pubKey,
		}
		//callHelloDesc := calldesc.NewClient(ctx, nil, sec)
		resHello, err := hubClient.JoinHello(ctx, reqHello)
		task.Assert(err)

		// Set computed encryption key for the client
		c.encKey = keyExch.ComputeKey(resHello.GetDhPubKey())
		node.UpdateNode(c.encKey)

		// Finish the join procedure
		//reqJoin := &hub.Join_Request{
		//	Tag:           c.nodeTags,
		//	AvailServices: nil,
		//	Flags:         nil,
		//	JoinSecret:    c.secret,
		//}
		//ctx = context.Background()
		//secPolicy := calldesc.NewSecurityPolicy(true, c.encKey)
		//desc := calldesc.NewClient(ctx, secPolicy)
		//desc.WithSessionFrom(resHello.)
		//resJoin, err := hubClient.Join(desc, reqJoin)
		//task.Assert(err)
		//_ = resJoin

		// Persist data of the newly joined node
		node.UpdateNode(c.encKey)
		task.Done()
	}
	return nil, run, nil
}