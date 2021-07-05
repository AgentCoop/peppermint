package join

import (
	"context"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/logger"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/node/model"
	"github.com/AgentCoop/peppermint/internal/security"
	"github.com/AgentCoop/peppermint/pkg"

	cc "github.com/AgentCoop/peppermint/internal/service/hub/grpc/client"
	//hh "github.com/AgentCoop/peppermint/internal/plugin/hub"
	//"github.com/AgentCoop/peppermint/internal/runtime"
)

func (c *joinContext) JoinTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		info := job.Logger(logger.Info)
		appNode := runtime.GlobalRegistry().App().(pkg.AppNode)
		ctx := context.Background()
		hubClient := j.GetValue().(cc.HubClient)

		// Generate a DH public key and exchange it with the hub server
		info("exchanging public keys with hub server...")
		keyExch, err := security.NewKeyExchange()
		task.Assert(err)

		pubKey := keyExch.GetPublicKey()
		reqHello := &hub.ClientHello_Request{DhPubKey: pubKey}
		resHello, err := hubClient.ClientHello(ctx, reqHello)
		task.Assert(err)

		// Set computed encryption key for the client
		c.encKey, err = keyExch.ComputeKey(resHello.GetDhPubKey())
		task.Assert(err)

		nodeDb := model.NewDb(appNode.Db())
		err = nodeDb.UpdateNode(appNode.Node().Id(), c.encKey)
		task.Assert(err)

		info("computed encryption key is %x...%x", c.encKey[0:1], c.encKey[len(c.encKey)-1:])
		//rt.NodeConfigurator().Refresh()

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

		info("join accepted")
		task.Done()
	}
	return nil, run, nil
}