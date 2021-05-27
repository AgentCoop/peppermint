package hub_test

import (
	"github.com/AgentCoop/peppermint/internal/app/node"
	cmd "github.com/AgentCoop/peppermint/internal/service/hub/service/client/join"
	"os"
	"testing"
	"time"

	_ "github.com/AgentCoop/peppermint/internal/service/hub"
)

var (
	serverAddr = "localhost:9911"
)

func TestJoinHello(t *testing.T) {
	os.Args = []string{"testapp", "run", "--hub-port=9911"}
	appJob := node.AppInitTest()
	appJob.Run()

	time.Sleep(100 * time.Millisecond)
	clientJob := cmd.JoinCmd(serverAddr, "secret")
	<-clientJob.Run()

	_, err := clientJob.GetInterruptedBy()
	if err != nil {
		t.Error(err)
	}
}
