package cmd

import (
	"github.com/AgentCoop/peppermint/internal/model/node"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

func DbMigrateCmd() {
	node.Migrate()
	runtime.GlobalRegistry().InvokeHooks(runtime.CmdMigrateDbHook)
}
