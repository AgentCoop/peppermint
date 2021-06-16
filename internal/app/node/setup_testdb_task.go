package node

import (
	job "github.com/AgentCoop/go-work"
	cmd "github.com/AgentCoop/peppermint/internal/app/node/cmd"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

func (app *app) SetupTestDbTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		// Fetch configurator data from DB and
		// and merge command-line options with fetched data
		//regServices := runtime.GlobalRegistry().Services()
		//for _, desc := range regServices {
		//	desc.Cfg.Fetch()
		//	desc.Cfg.MergeCliOptions(app.CliParser())
		//}
	}
	run := func(task job.Task) {
		db := runtime.GlobalRegistry().Db()
		parser := app.CliParser()
		cmd.DbCreateCmd(db, parser, true)
		cmd.BootstrapCmd("", nil)
		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}


