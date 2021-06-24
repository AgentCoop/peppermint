package main

import (
	"github.com/AgentCoop/peppermint/internal/service/test/grpc/client"
	"github.com/jessevdk/go-flags"
	"net"
)

func main() {
	parser := flags.NewParser(&options, flags.IgnoreUnknown)
	_, err := parser.Parse()
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", options.Service)
	if err != nil {
		panic(err)
	}
	cc := client.NewClient(addr)
	cmdName := parser.Active.Name
	var cmdOptions interface{}
	switch cmdName {
	case client.CMD_NAME_PING:
		cmdOptions = options.CmdPingOptions
	}
	ctx := client.NewCmdContext(cmdOptions, options.Count)
	cmdJob := client.NewJob(cmdName, ctx, cc)
	// Execute command
	<-cmdJob.Run()
	// Handle error
	_, jobErr := cmdJob.GetInterruptedBy()
	if jobErr != nil {
		panic(jobErr)
	}
}
