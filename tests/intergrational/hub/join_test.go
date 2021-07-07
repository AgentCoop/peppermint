package hub_test

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/app/node"
	_ "github.com/AgentCoop/peppermint/internal/service/hub"
	"github.com/AgentCoop/peppermint/internal/service/hub/grpc/client"
	"github.com/AgentCoop/peppermint/internal/service/hub/grpc/client/join"
	"os"
	"testing"
	"time"
)

var (
	serverAddr = "localhost:9911"
)

func TestJoinHello(t *testing.T) {
	os.Args = []string{"testapp", "bootstrap", "--create-db", "--force"}
	app := node.NewApp()
	<-app.Job().Run()

	os.Args = []string{"testapp", "run", "--hub-port=12001", "-n=1"}
	appJob := node.NewApp()
	appJob.Job().Run()

	time.Sleep(50 * time.Millisecond)

	//addr, _ := net.ResolveTCPAddr("tcp", "localhost:9911")
	//time.Sleep(time.Millisecond)
	hubClient := client.NewClient("localhost")

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
