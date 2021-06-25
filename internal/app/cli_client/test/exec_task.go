package test

import (
	"context"
	job "github.com/AgentCoop/go-work"
	"time"
	tt "github.com/AgentCoop/peppermint/internal/api/peppermint/service/frontoffice/test"
)

func (app *appTest) CmdExecutorTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		callFn := <- app.execCmdChan
		totalFailedCount := 0
		for i := 0; i < int(app.callParams.workers); i++ {
			go func() {
				for j := 0; j < int(app.callParams.repeat); j++ {
					ctx := context.Background()
					t := app.callParams.timeout
					if t > 0 {
						dealine := time.Now().Add(time.Duration(t) * time.Millisecond)
						ctx, _ = context.WithDeadline(ctx, dealine)
					}
					params := app.callParams
					req := &tt.Request{
						RsDelay:                params.rsDelay,
						RsDelayJitter:          params.rsDelayJitter,
						RsBulkData:             nil,
						BulkData:               nil,
						CallFailureProbability: 0,
						Token:                  params.token,
					}
					_, err := callFn(ctx, req)
					if err != nil {
						totalFailedCount++
					}
				}
			}()
		}
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

