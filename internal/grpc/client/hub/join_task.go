package hub

import (
	job "github.com/AgentCoop/go-work"
)

func (c *hubClient) JoinTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

