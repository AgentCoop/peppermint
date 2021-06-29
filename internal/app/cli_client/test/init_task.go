package test

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/app"
)

func (appTest *appTest) InitTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		app.AppInit(appTest, task)
		task.Done()
	}
	return nil, run, nil
}
