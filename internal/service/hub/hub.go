
package hub

import (
	"fmt"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/cliparser"
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
	reg.RegisterParserCmdHook(cliparser.CREATEDB_CMD_NAME, hub.createDb)
}

func (h *hubService) createDb(data interface{}) {
	fmt.Printf("time to sleep, Andrew!\n")
}



