package service

import (
	job "github.com/AgentCoop/go-work"
)

func (svc *baseService) InitTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		task.Done()
	}
	return init, nil, nil
}
