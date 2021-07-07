package node

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
	cmd "github.com/AgentCoop/peppermint/internal/app/node/cmd"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/utils"
)

func (app *appNode) ParserTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		if cmd.Options.NodeId == 0 {
			cmd.Options.NodeId = 1
		}
	}
	run := func(task job.Task) {
		rt := runtime.GlobalRegistry().Runtime()
		parser := rt.CliParser()
		cmdName, _ := parser.CurrentCmd()
		switch cmdName {
		case cmd.CMD_NAME_BOOTSTRAP:
			opts, err := parser.GetCmdOptions(cmdName)
			task.Assert(err)
			v := opts.(cmd.Bootstrap)
			err = cmd.BootstrapCmd(v.CreateDb, v.Force, v.IdFromInterface, v.Tags)
			task.Assert(err)
			task.FinishJob()

		case cmd.CMD_NAME_VERSION:
			opts, err := parser.GetCmdOptions(cmdName)
			task.Assert(err)
			verbose := opts.(cmd.Version).Verbose
			cmd.ShowVersion(verbose)

		case cmd.CMD_NAME_DB_MIGRATE:
			// opts, err := parser.GetCmdOptions(cmdName)
			// task.Assert(err)
			// v := opts.(cmd.DbMigrate)
			cmd.DbMigrateCmd()

		case cmd.CMD_NAME_RUN:
			err := cmd.RunCmd()
			task.Assert(err)

		case cmd.CMD_NAME_JOIN:
			secret, err := utils.ReadPassword("Enter join secret")
			task.Assert(err)
			fmt.Println("")
			opts, err := parser.GetCmdOptions(cmdName)
			task.Assert(err)
			v := opts.(cmd.Join)
			err = cmd.JoinCmd(secret, v.Tags, v.Hub)
			task.Assert(err)
		default:
			fmt.Printf("no command\n")
		}
		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}
