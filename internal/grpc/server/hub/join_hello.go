package hub

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	"context"
	"github.com/AgentCoop/peppermint/internal/grpc"
	data "github.com/AgentCoop/peppermint/internal/grpc/data/hub/join"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	join "github.com/AgentCoop/peppermint/internal/service/hub/server/join"
	"github.com/AgentCoop/peppermint/internal/utils"
)

func (s *hubGrpcServer) JoinHello(ctx context.Context, originalReq *hub.JoinHello_Request) (*hub.Response, error) {
	req := data.NewJoinHello(ctx, originalReq)
	_ = req.Validate()

	j := job.NewJob(nil)
	sessId := server.StartNewSession(j)

	joinCtx := join.NewJoinContext()
	j.AddTask(joinCtx.JoinHelloTask)
	j.AddTask(joinCtx.JoinTask)
	j.Run()

	joinCtx.JoinHelloReqCh <- req
	res := <-joinCtx.JoinHelloRespCh

	res.AddMetaValue(grpc.META_FIELD_SESSION_ID, utils.IntToHex(sessId, 16))

	return res.ToGrpcResponse().(*hub.Response), nil
}
