package cmd

import (
	"github.com/AgentCoop/peppermint/internal/model/node"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

func DbCreateCmd(force bool) {
	if force {
		node.DropTables()
	}
	node.CreateTables()
	runtime.GlobalRegistry().InvokeHooks(runtime.CmdCreateDbHook, force)
}
