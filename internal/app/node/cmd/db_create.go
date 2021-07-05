package cmd

import (
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/node/model"
)

func DbCreateCmd(force bool) {
	app := runtime.GlobalRegistry().App()
	appDb := model.NewDb(app.Db())
	if force {
		appDb.DropTables()
	}
	appDb.CreateTables()
	runtime.GlobalRegistry().InvokeHooks(runtime.CmdCreateDbHook, force)
}
