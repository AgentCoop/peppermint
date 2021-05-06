package hub

import (
	"github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	g "github.com/AgentCoop/peppermint/internal/grpc/server"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	hub.UnimplementedHubServer
	grpc *grpc.Server
	lis net.Listener
}

func NewServer() *server {
	s := &server{}
	s.grpc = g.DefaultGrpcServer()
	return s
}

func (s *server) StartServerTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		lis, err := net.Listen("tcp", "localhost:9000")
		task.Assert(err)

		s.lis = lis
		hub.RegisterHubServer(s.grpc, s)
	}
	run := func (task job.Task) {
		s.grpc.Serve(s.lis)
		task.Done()
	}
	return init, run, nil
}
