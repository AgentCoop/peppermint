package cmd

import (
	"fmt"
	"github.com/AgentCoop/peppermint/cmd"
	"github.com/AgentCoop/peppermint/internal/db"
	"github.com/AgentCoop/peppermint/internal/runtime/deps"

	//"github.com/AgentCoop/peppermint/internal/model/hub"
	//"github.com/AgentCoop/peppermint/internal/model/webproxy"
	"github.com/AgentCoop/peppermint/internal/runtime"
	//"github.com/AgentCoop/peppermint/internal/service/webproxy/model"
	"gorm.io/gorm"
)

func dropAllTables(migrator gorm.Migrator) {
	//migrator.DropTable(
	//	model.WebProxyConfig{},
	//	model.HubConfig{},
	//)
}

func DbMigrateCmd(db db.Db, parser deps.CliParser, drop bool) {
	h := db.Handle()
	if drop {
		fmt.Printf("-> dropping all tables before migration\n")
		dropAllTables(h.Migrator())
	}
	// Invoke all registered database migration hooks
	opts, _ := parser.GetCmdOptions(cmd.CMD_NAME_DB_MIGRATE)
	for _, hook := range runtime.GlobalRegistry().LookupParserCmdHook(cmd.CMD_NAME_DB_MIGRATE) {
		hook(opts)
	}
}
