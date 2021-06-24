package test

import (
	job "github.com/AgentCoop/go-work"
	api "github.com/AgentCoop/peppermint/internal/api/peppermint/service/frontoffice/test"
	g "github.com/AgentCoop/peppermint/internal/service/test/grpc/server"
	"github.com/AgentCoop/peppermint/internal/service/test/logger"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/service"
	"net"
	"strconv"
)

var (
	Name = api.Test_ServiceDesc.ServiceName
)

type testService struct {
	runtime.Service
}

func init() {
	test := new(testService)
	reg := runtime.GlobalRegistry()
	reg.RegisterHook(runtime.ServiceInitHook, func(args...interface{}) {
		test.Init()
	})
	reg.RegisterHook(runtime.CmdCreateDbHook, func(args...interface{}) {
		test.createDd(args...)
	})
}

func (test *testService) Init() (runtime.Service, error) {
	rt := runtime.GlobalRegistry().Runtime()
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:"+strconv.Itoa(12099))
	// Create network server and service policy
	srv := g.NewServer(Name, addr)
	srv.WithStdoutLogger(job.Logger(logger.Info))
	policy := service.NewServicePolicy(srv.FullName(), srv.Methods())
	test.Service = service.NewBaseService(srv, nil, nil, policy)
	rt.RegisterService(Name, test)
	return test, nil
}

func (test *testService) createDd(args...interface{}) {
	//force := args[0].(bool)
	//if force {
	//	model.DropTables()
	//}
	//model.CreateTables()
}
