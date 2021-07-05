package node

import (
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/pkg"
	"os"
)

type appNode struct {
	pkg.App
	sigChan chan os.Signal
	node    pkg.Node
}

func (app *appNode) Node() pkg.Node {
	return app.node
}

func (app *appNode) reloadServiceConfig() {
	rt := runtime.GlobalRegistry().Runtime()
	for _, svc := range rt.Services() {
		svc.ReloadConfig(app.node.Id())
	}
}
