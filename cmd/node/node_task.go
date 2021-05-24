package main

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/cmd"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

func (app *app) NodeTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		// Merge command-line options with DB settings
		regServices := runtime.GlobalRegistry().Services()
		for _, item := range regServices {
			item.Cfg.Fetch()
			item.Cfg.MergeCliOptions(app.CliParser())
		}
	}
	run := func(task job.Task) {
		cmdName, _ := app.CliParser().CurrentCmd()
		switch cmdName {
		case cmd.CMD_NAME_VERSION:
			opts, err := app.CliParser().GetCmdOptions(cmdName)
			task.Assert(err)
			verbose := opts.(Version).Verbose
			showVersion(verbose)

		case cmd.CMD_NAME_DB_MIGRATE:
			fmt.Printf("db migrate\n")

		case cmd.CMD_NAME_RUN:
			serviceJob := job.NewJob(nil)
			regServices := runtime.GlobalRegistry().Services()
			for _, info := range regServices {
				service := info.Initializer()
				serviceJob.AddTask(service.StartTask)
			}
			<-serviceJob.Run()

		case cmd.CMD_NAME_JOIN:
			fmt.Printf("join\n")
		default:
			fmt.Printf("no command\n")
		}
		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}


