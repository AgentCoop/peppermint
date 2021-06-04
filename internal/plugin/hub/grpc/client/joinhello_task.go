package client

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/grpc/client"
	"github.com/AgentCoop/peppermint/internal/grpc/communicator"
	data "github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/data/client/join"
	//"github.com/AgentCoop/peppermint/internal/service/hub/client/join"

	//"github.com/AgentCoop/peppermint/internal/api/peppermint/service"
)

func (c *hubClient) JoinHelloTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		comm := j.GetValue().(session.ClientGrpcServiceCommunicator)
		dataBag := comm.GrpcRx(0)
		task.Assert(dataBag)
		callDesc := dataBag.(client.ClientCallDescriptor)

		req := dataBag.(client.ClientCallDescriptor).GetRequest()
		origResp, err := c.grpcHandle.JoinHello(callDesc, req.ToGrpcRequest().(*hub.JoinHello_Request))
		task.Assert(err)

		data.NewJoinHelloResponse(callDesc, origResp)
		comm.GrpcTx(0)

		task.Done()
	}
	return nil, run, nil
}