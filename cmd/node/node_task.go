package main

import (
	job "github.com/AgentCoop/go-work"
)

func (app *app) NodeTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		cmdName, _ := app.CliParser().CurrentCmd()
		app.handleCmdOptions(cmdName, task)
		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}


