package server

import (
	"context"
	job "github.com/AgentCoop/go-work"
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	data "github.com/AgentCoop/peppermint/internal/service/hub/grpc/data/server/join"
	srv "github.com/AgentCoop/peppermint/internal/grpc/server"
	//srv "github.com/AgentCoop/peppermint/internal/grpc/server/hub"
	join "github.com/AgentCoop/peppermint/internal/service/hub/service/server/join"
)

func (s *hubServer) JoinHello(ctx context.Context, originalReq *msg.JoinHello_Request) (*msg.JoinHello_Response, error) {
	pair := ctx.(srv.RequestResponsePair)
	req := data.NewJoinHello(pair, originalReq)
	_ = req.Validate()

	j := job.NewJob(nil)
	sessId := srv.StartNewSession(j)
	_ = sessId

	joinCtx := join.NewJoinContext()
	j.AddTask(joinCtx.JoinHelloTask)
	j.AddTask(joinCtx.JoinTask)
	j.Run()

	// Dispatch request-response pair to JoinHelloTask started above
	joinCtx.ReqChan[0] <- pair
	<-joinCtx.ResChan[0]

	res := pair.GetResponse()
	res.SetSessionId(3)

	return res.ToGrpcResponse().(*msg.JoinHello_Response), nil
}
