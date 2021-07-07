package node

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/app"
	"github.com/AgentCoop/peppermint/internal/app/node/cmd"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/node/model"
)

func (appNode *appNode) InitTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		err := app.AppInit(appNode)
		task.Assert(err)
	}
	run := func(task job.Task) {
		rt := runtime.GlobalRegistry().Runtime()
		parser := rt.CliParser()
		db, err := app.InitDb(appNode,"node.db")
		task.Assert(err)
		appNode.WithDb(db)
		// Init node configurator
		cmdName, _ := parser.CurrentCmd()
		switch cmdName {
		case cmd.CMD_NAME_RUN:
			nodeDb := model.NewDb(db)
			node, err := nodeDb.NewNode(cmd.Options.NodeId)
			task.Assert(err)
			appNode.node = node
			// Service db
			appNode.InitServiceDb(node)
			// Initialize services
			err = runtime.GlobalRegistry().InvokeHooks(runtime.ServiceInitHook)
			task.Assert(err)
		}
		task.Done()
	}
	return init, run, nil
}
