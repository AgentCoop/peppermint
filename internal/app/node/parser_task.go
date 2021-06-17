package node

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
	cmd "github.com/AgentCoop/peppermint/internal/app/node/cmd"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/utils"
)

func (app *app) ParserTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		parser := app.CliParser()
		cmdName, _ := parser.CurrentCmd()
		switch cmdName {
		case cmd.CMD_NAME_DB_CREATE, cmd.CMD_NAME_BOOTSTRAP:
		default:
			err := app.NodeConfigurator().Fetch()
			task.Assert(err)
		}
		// Services initialization
		runtime.GlobalRegistry().InvokeHooks(runtime.ServiceInitHook)
	}
	run := func(task job.Task) {
		db := runtime.GlobalRegistry().Db()
		parser := app.CliParser()
		cmdName, _ := parser.CurrentCmd()
		switch cmdName {
		case cmd.CMD_NAME_BOOTSTRAP:
			opts, err := parser.GetCmdOptions(cmdName)
			task.Assert(err)
			v := opts.(cmd.Bootstrap)
			cmd.BootstrapCmd(v.IdFromInterface, v.Tags)

		case cmd.CMD_NAME_VERSION:
			opts, err := parser.GetCmdOptions(cmdName)
			task.Assert(err)
			verbose := opts.(cmd.Version).Verbose
			cmd.ShowVersion(verbose)

		case cmd.CMD_NAME_DB_MIGRATE:
			opts, err := parser.GetCmdOptions(cmdName)
			task.Assert(err)
			v := opts.(cmd.DbMigrate)
			cmd.DbMigrateCmd(db, parser, v.Drop)

		case cmd.CMD_NAME_DB_CREATE:
			opts, err := parser.GetCmdOptions(cmdName)
			task.Assert(err)
			v := opts.(cmd.DbCreate)
			cmd.DbCreateCmd(v.Force)

		case cmd.CMD_NAME_RUN:
			cmd.RunCmd()

		case cmd.CMD_NAME_JOIN:
			secret, err := utils.ReadPassword("Enter join secret")
			task.Assert(err)
			opts, err := parser.GetCmdOptions(cmdName)
			task.Assert(err)
			v := opts.(cmd.Join)
			cmd.JoinCmd(secret, v.Tags, v.Hub)
		default:
			fmt.Printf("no command\n")
		}
		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}


