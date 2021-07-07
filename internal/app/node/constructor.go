package node

import (
	app "github.com/AgentCoop/peppermint/internal/app"
	"github.com/AgentCoop/peppermint/internal/app/node/cmd"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"os"
	"os/signal"
	"syscall"
)

func NewApp() *appNode {
	appNode := new(appNode)
	// Listen to some system signals
	appNode.sigChan = make(chan os.Signal, 1)
	signal.Notify(appNode.sigChan, syscall.SIGHUP)
	appNode.App = app.NewApp(&cmd.Options)
	// Main job
	appNode.Job().AddOneshotTask(appNode.InitTask)
	appNode.Job().AddTask(appNode.ParserTask)
	appNode.Job().AddTask(appNode.SigTask)
	// Add created application to the global registry
	runtime.GlobalRegistry().SetApp(appNode)
	return appNode
}
