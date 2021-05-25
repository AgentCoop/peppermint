package join

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/runtime/config"
	data "github.com/AgentCoop/peppermint/internal/service/hub/grpc/data/server/join"
)

func (ctx *joinCtx) JoinTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		pair := <-ctx.reqChan[1]
		task.AssertNotNil(pair)
		req := pair.GetRequest()

		dataBag := req.(data.Join_DataBag)
		secret, tags := dataBag.Secret(), dataBag.Tags()

		cfg := pair.GetConfigurator().(config.HubConfigurator)
		if cfg.Secret() != secret {
			panic("wrong secret word")
		}

		_, _ = secret, tags
	}
	return nil, run, nil
}

