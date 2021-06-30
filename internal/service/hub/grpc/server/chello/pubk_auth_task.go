package chello

import job "github.com/AgentCoop/go-work"

func (ctx *clientHelloCtx) PubKeyAuthTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {

	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

