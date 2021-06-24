package app

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/utils"
)

func (app *app) InitTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		rt := runtime.GlobalRegistry().Runtime()
		err := rt.CliParser().Run()
		task.Assert(err)

		err = utils.FS_FileOrDirExists(app.appDir)
		task.Assert(err)

		task.Done()
	}
	return nil, run, nil
}
