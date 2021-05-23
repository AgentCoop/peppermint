package hub

import job "github.com/AgentCoop/go-work"

func (s *hubService) StartTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {

	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}
