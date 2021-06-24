package client

import (
	"context"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/frontoffice/test"
)

func (ctx *cmdContext) PingTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		testClient := j.GetValue().(TestClient)
		req := &test.Ping_Request{Msg: "Ping"}
		for i := 0; i < ctx.count; i++ {
			res, err := testClient.Ping(context.Background(), req)
			_ = res
			task.Assert(err)
		}
		task.Done()
	}
	return nil, run, nil
}