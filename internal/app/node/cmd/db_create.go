package cmd

import (
	"github.com/AgentCoop/peppermint/cmd"
	"github.com/AgentCoop/peppermint/internal/db"
	"github.com/AgentCoop/peppermint/internal/model/node"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

func DbCreateCmd(db db.Db, parser runtime.CliParser, force bool) {
	h := db.Handle()
	m := h.Migrator()

	if force { m.DropTable(node.Tables...) }
	h.Migrator().CreateTable(node.Tables...)

	// Invoke all registered database migration hooks
	opts, _ := parser.GetCmdOptions(cmd.CMD_NAME_DB_MIGRATE)
	for _, hook := range runtime.GlobalRegistry().LookupParserCmdHook(cmd.CMD_NAME_DB_MIGRATE) {
		hook(opts)
	}
}
