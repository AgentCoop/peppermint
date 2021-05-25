package server

import (
	"context"
	job "github.com/AgentCoop/go-work"
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	srv "github.com/AgentCoop/peppermint/internal/grpc/server"
	"github.com/AgentCoop/peppermint/internal/runtime"
	data "github.com/AgentCoop/peppermint/internal/service/hub/grpc/data/server/join"
	"github.com/AgentCoop/peppermint/internal/service/hub/service/server/join"
)

func (s *hubServer) JoinHello(ctx context.Context, originalReq *msg.JoinHello_Request) (*msg.JoinHello_Response, error) {
	pair := ctx.(srv.RequestResponsePair)
	req := data.NewJoinHello(pair, originalReq)
	_ = req.Validate()

	joinCtx := join.NewJoinContext()
	j := job.NewJob(joinCtx)
	j.AddTask(joinCtx.JoinHelloTask)
	j.AddTask(joinCtx.JoinTask)
	j.Run()

	sessId := runtime.GlobalRegistry().GrpcSession().New(j, 3600)
	// Dispatch request-response pair to JoinHelloTask started above
	joinCtx.ReqChan()[0] <- pair
	<-joinCtx.ResChan()[0]

	res := pair.GetResponse()
	res.SetSessionId(sessId)

	return res.ToGrpcResponse().(*msg.JoinHello_Response), nil
}
