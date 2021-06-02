package join

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	data "github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/data/server/join"
	"github.com/AgentCoop/peppermint/internal/plugin/hub/model"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ctx *joinCtx) JoinTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		comm := j.GetValue().(runtime.GrpcServiceCommunicator)
		v := comm.ServiceRx(1)
		task.AssertNotNil(v)
		callDesc := v.(server.GrpcCallDescriptor)
		req := callDesc.GetRequest()

		dataBag := req.(data.Join_DataBag)
		secret, tags := dataBag.Secret(), dataBag.Tags()

		cfg := callDesc.GetConfigurator().(config.HubConfigurator)
		if cfg.Secret() != secret {
			comm.ServiceTx(1, status.Error(codes.PermissionDenied, "Invalid join secret"))
			task.Done()
			return
		}

		joinedId := 1//utils.Rand_UniqueId()
		data.NewJoinResponse(callDesc, 0)

		// Persist node data
		_ = tags
		newNode := &model.HubJoinedNode{
			EncKey: ctx.encKey,
			NodeId: uint64(joinedId),
			Tags:   []model.HubNodeTag{{Name: "foo222"}},
		}

		db := runtime.GlobalRegistry().Db().Handle()
		db.Create(newNode)
		task.Assert(db.Error)
		lastError := db.Error
		_ = lastError

		comm.ServiceTx(1, callDesc)
		task.Done()
	}
	return nil, run, nil
}

