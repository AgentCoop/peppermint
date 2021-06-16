package node

import (
	job "github.com/AgentCoop/go-work"
	cmd "github.com/AgentCoop/peppermint/internal/app/node/cmd"
)

func (app *app) SetupTestDbTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		cmd.DbCreateCmd(true)
		cmd.BootstrapCmd("", nil)
		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}


