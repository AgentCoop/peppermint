package hub

import (
	"context"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	//"github.com/AgentCoop/peppermint/internal/api/peppermint/service"
)

func (c *hubClient) JoinHello(dhPubKey []byte) []byte {
	grpc := hub.NewHubClient(c.Conn)
	r := &hub.JoinHello_Request{DhPubKey: dhPubKey}
	res, err := grpc.JoinHello(context.Background(), r)
	c.Task.Assert(err)
	return res.DhPubKey
}

func (c *hubClient) JoinHelloTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {

	}
	return nil, run, nil
}