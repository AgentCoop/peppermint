package main

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/cmd"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/utils"
)

func (app *app) NodeTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		// Merge command-line options with DB settings
		regServices := runtime.GlobalRegistry().Services()
		for _, desc := range regServices {
			desc.Cfg.Fetch()
			desc.Cfg.MergeCliOptions(app.CliParser())
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
			for _, desc := range regServices {
				service := desc.Initializer()
				serviceJob.AddTask(service.StartTask)
			}
			<-serviceJob.Run()

		case cmd.CMD_NAME_JOIN:
			secret, err := utils.ReadPassword("Enter join secret")
			task.Assert(err)
			opts, err := app.CliParser().GetCmdOptions(cmdName)
			task.Assert(err)
			v := opts.(Join)
			joinCmd(secret, v.Tags, v.Hub)
		default:
			fmt.Printf("no command\n")
		}
		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}


