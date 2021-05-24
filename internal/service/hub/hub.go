
package hub

import (
	"fmt"
	"github.com/AgentCoop/peppermint/cmd"
	model "github.com/AgentCoop/peppermint/internal/model/hub"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

const (
	Name = "Hub"
)

type hubService struct {

}

func init() {
	hub := &hubService{}
	reg := runtime.GlobalRegistry()
	reg.RegisterService(Name, hub)
	reg.RegisterParserCmdHook(cmd.CMD_NAME_DB_MIGRATE, hub.migrateDb)
}

func (h *hubService) migrateDb(data interface{}) {
	db := runtime.GlobalRegistry().Db()
	gorm := db.Handle()
	_ = gorm.AutoMigrate(&model.JoinedNode{})
	fmt.Printf("time to sleep, Andrew!\n")
}



