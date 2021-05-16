package hub

import (
	"context"
	job "github.com/AgentCoop/go-work"
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	data "github.com/AgentCoop/peppermint/internal/grpc/data/hub/join"
	srv "github.com/AgentCoop/peppermint/internal/grpc/server"
	//srv "github.com/AgentCoop/peppermint/internal/grpc/server/hub"
	join "github.com/AgentCoop/peppermint/internal/service/hub/server/join"
)

func (s *hubServer) JoinHello(ctx context.Context, originalReq *msg.JoinHello_Request) (*msg.JoinHello_Response, error) {
	req := data.NewJoinHello(ctx, originalReq)
	_ = req.Validate()

	j := job.NewJob(nil)
	sessId := srv.StartNewSession(j)

	joinCtx := join.NewJoinContext()
	j.AddTask(joinCtx.JoinHelloTask)
	j.AddTask(joinCtx.JoinTask)
	j.Run()

	joinCtx.JoinHelloReqCh <- req
	res := <-joinCtx.JoinHelloRespCh

	resHeader := res.(srv.ResponseHeader)
	resHeader.SetSessionId(sessId)

	return res.ToGrpcResponse().(*msg.JoinHello_Response), nil
}
