package cmd

import (
	"github.com/AgentCoop/peppermint/internal/model/node"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

func DbCreateCmd(force bool) {
	node.CreateTables()
	runtime.GlobalRegistry().InvokeHooks(runtime.CmdCreateDbHook)
}
