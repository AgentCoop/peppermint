package cmd

import (
	"github.com/AgentCoop/peppermint/internal/db"
	"github.com/AgentCoop/peppermint/internal/model/node"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

func DbCreateCmd(db db.Db, parser runtime.CliParser, force bool) {
	node.CreateTables()
	runtime.GlobalRegistry().InvokeHooks(runtime.CmdCreateDbHook)
}
