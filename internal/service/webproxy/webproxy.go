package webproxy

import (
	"github.com/AgentCoop/peppermint/cmd"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

const (
	Name = "WebProxy"
)

type hubService struct {

}

func init() {
	hub := &hubService{}
	reg := runtime.GlobalRegistry()
	reg.RegisterService(Name, hub)
	//reg.RegisterParserCmdHook(cmd.CMD_NAME_DB_MIGRATE, hub.migrateDb)
}
