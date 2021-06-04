package hub_test

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/app/node"
	"github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/client"
	"github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/client/join"
	"net"
	"os"
	"testing"
	"time"
)

var (
	serverAddr = "localhost:9911"
)

func TestJoinHello(t *testing.T) {
	os.Args = []string{"testapp", "run", "--hub-port=9911"}
	appJob := node.AppInitTest()
	appJob.Run()

	time.Sleep(100 * time.Millisecond)

	addr, _ := net.ResolveTCPAddr("tcp", "localhost:9991")
	hubClient := client.NewClient(addr)

	clientJob := job.NewJob(nil)
	joinCtx := join.NewJoinContext("secret", []string{"my-test-machine", "linux"})
	clientJob.AddOneshotTask(hubClient.ConnectTask)
	clientJob.AddTask(joinCtx.JoinTask)
	<-clientJob.Run()

	_, err := clientJob.GetInterruptedBy()
	if err != nil {
		t.Error(err)
	}
}
