package main

import (
	"fmt"
	"github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

const (
	CLI_CMD_ARG_JOIN = "join"
	CLI_CMD_ARG_CREATEDB = "createdb"
)

var (
	mainJob job.Job = job.NewJob(nil)

	//
	cliParser runtime.CliParser
)

func addJoinCommand() {
	cliParser.AddCommand(CLI_CMD_ARG_JOIN, &options.JoinCmd, "", "")
}

func addCreateDb() {
	cliParser.AddCommand(CLI_CMD_ARG_CREATEDB, &options.CreateDb, "", "")
}

func main() {
	r := runtime.NewRuntime(runtime.NewCliParser(&options))
	mainJob.SetValue(r)

	//mainJob = job.NewJob(nil)
	//cliParser = runtime.NewCliParser(&options)
	//clie
	fmt.Printf("main\n")
}
