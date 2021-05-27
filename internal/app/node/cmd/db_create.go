package cmd

import (
	"github.com/AgentCoop/peppermint/cmd"
	"github.com/AgentCoop/peppermint/internal/db"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

func DbCreateCmd(db db.Db, parser runtime.CliParser, force bool) {
	h := db.Handle()
	_ = h
	if force {

	}
	// Invoke all registered database migration hooks
	opts, _ := parser.GetCmdOptions(cmd.CMD_NAME_DB_MIGRATE)
	for _, hook := range runtime.GlobalRegistry().LookupParserCmdHook(cmd.CMD_NAME_DB_MIGRATE) {
		hook(opts)
	}
}
