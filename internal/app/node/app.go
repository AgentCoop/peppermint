package node

import (
	"github.com/AgentCoop/peppermint/pkg"
)

type appNode struct {
	pkg.App
	node pkg.Node
}

//func (app *appNode) Db() pkg.Db {
//	app.App.Db()
//}

func (app *appNode) Node() pkg.Node {
	return app.node
}
