package hub

import (
	"context"
	job "github.com/AgentCoop/go-work"
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	"github.com/AgentCoop/peppermint/internal/grpc"
	data "github.com/AgentCoop/peppermint/internal/grpc/data/hub/join"
	srv "github.com/AgentCoop/peppermint/internal/grpc/server"
	//srv "github.com/AgentCoop/peppermint/internal/grpc/server/hub"
	join "github.com/AgentCoop/peppermint/internal/service/hub/server/join"
	"github.com/AgentCoop/peppermint/internal/utils"
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

	res.AddMetaValue(grpc.META_FIELD_SESSION_ID, utils.IntToHex(sessId, 16))

	return res.ToGrpcResponse().(*msg.JoinHello_Response), nil
}
