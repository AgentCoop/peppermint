package node

import (
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/app"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/utils"
	"github.com/AgentCoop/peppermint/pkg"
	"os"
)

type appNode struct {
	pkg.App
	sigChan chan os.Signal
	node    pkg.Node
	svcDb   pkg.Db
}

func (app *appNode) Node() pkg.Node {
	return app.node
}

func (app *appNode) ServiceDb() pkg.Db {
	return app.svcDb
}

func (appNode *appNode) InitServiceDb(node pkg.Node) error {
	basename := utils.Conv_IntToHex(node.ExternalId(), internal.NodeId(0).Size())
	basename += "_svc.db"
	svcDb, err := app.InitDb(appNode, basename)
	if err != nil {
		return err
	}
	appNode.svcDb = svcDb
	return nil
}

func (app *appNode) reloadServiceConfig() {
	rt := runtime.GlobalRegistry().Runtime()
	for _, svc := range rt.Services() {
		svc.ReloadConfig(app.node.Id())
	}
}
