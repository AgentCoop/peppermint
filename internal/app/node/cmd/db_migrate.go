package cmd

import (
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/node/model"
)

func DbMigrateCmd() {
	model.Migrate()
	runtime.GlobalRegistry().InvokeHooks(runtime.CmdMigrateDbHook)
}
