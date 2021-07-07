package cmd

import (
	"crypto/sha256"
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/node/model"
	"github.com/AgentCoop/peppermint/internal/utils"
	"github.com/AgentCoop/peppermint/pkg"
)

func BootstrapCmd(createDb bool, force bool, IdFromNic string, tags []string) error {
	var nodeId internal.NodeId
	app := runtime.GlobalRegistry().App().(pkg.AppNode)
	switch {
	case len(IdFromNic) != 0: // Generate node ID from the given network interface
		macAddr, err := utils.Net_MacAdrr(IdFromNic)
		if err != nil {
			return err
		}
		hash := sha256.Sum256(macAddr)
		nodeId.FromByteArray(hash[:])
	default: // Random node ID
		nodeId.Rand()
	}
	appDb := model.NewDb(app.Db())
	// Create common database
	if createDb {
		if force {
			appDb.DropTables()
		} else {
			// Check if database was already created
			_, err := appDb.FetchById(1)
			if err != nil {
				panic("db already exists")
			}
		}
		appDb.CreateTables()
	}
	node, err := appDb.CreateNode(nodeId, tags)
	if err != nil {
		return err
	}
	if err := app.InitServiceDb(node); err != nil {
		return err
	}
	err = runtime.GlobalRegistry().InvokeHooks(runtime.CmdCreateDbHook, true, app.ServiceDb())
	return err
}
