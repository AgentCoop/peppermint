// +build with_hub

package main

import (
//	"github.com/AgentCoop/peppermint/internal/runtime"
	"fmt"
	"github.com/AgentCoop/peppermint/internal/service/hub"
)

func init() {
	fmt.Printf("with hub\n")
	s := hub.NewService()
	_ = s
	//runtime.AddCliArgHook(runtime.CLI_CMD_ARG_CREATEDB, func() {
	//	mainJob.AddTask(s.CreateDbTask())
	//})
}
