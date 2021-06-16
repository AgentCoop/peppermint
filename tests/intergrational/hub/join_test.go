package hub_test

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/app/node"
	_ "github.com/AgentCoop/peppermint/internal/plugin/hub"
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

	addr, _ := net.ResolveTCPAddr("tcp", "localhost:9911")
	time.Sleep(time.Millisecond)
	hubClient := client.NewClient(addr)

	clientJob := job.NewJob(hubClient)
	joinCtx := join.NewJoinContext("secret", []string{"my-test-machine", "linux"})
	clientJob.AddOneshotTask(hubClient.ConnectTask)
	clientJob.AddTask(joinCtx.JoinTask)
	<-clientJob.Run()

	_, err := clientJob.GetInterruptedBy()
	if err != nil {
		t.Error(err)
	}
}
