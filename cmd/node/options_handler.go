package main

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/cmd"
)

func (app *app) handleCmdOptions(cmdName string, task job.Task) {
	switch cmdName {
	case cmd.CMD_NAME_VERSION:
		opts, err := app.CliParser().GetCmdOptions(cmdName)
		task.Assert(err)
		verbose := opts.(Version).Verbose
		showVersion(verbose)
	case cmd.CMD_NAME_DB_MIGRATE:
		fmt.Printf("db migrate\n")
	case cmd.CMD_NAME_RUN:
		fmt.Printf("run\n")
	case cmd.CMD_NAME_JOIN:
		fmt.Printf("join\n")
	default:
		fmt.Printf("no command\n")
	}
}
