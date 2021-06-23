package proxy

import (
	job "github.com/AgentCoop/go-work"
	g "github.com/AgentCoop/peppermint/internal/grpc"
)

func NewProxyLink(upstream *proxyStream, downstream *proxyStream) (job.Job, *proxyConn) {
	pconn := &proxyConn{
		upstream:   upstream,
		downstream: downstream,
		downChan:   make(chan g.CodecPacket, 1),
		upChan:     make(chan g.CodecPacket, 1),
	}
	pjob := job.NewJob(nil)
	pjob.AddTask(pconn.readUpstreamTask)
	pjob.AddTask(pconn.readDownstreamTask)
	pjob.AddTask(pconn.writeStreamTask)
	return pjob, pconn
}

func NewProxyLinkFromUpClient(upClient g.BaseClient, downstream *proxyStream) (job.Job, *proxyConn) {
	pjob, pconn := NewProxyLink(nil, downstream)
	pjob.AddOneshotTask(upClient.ConnectTask)
	pjob.AddTask(pconn.initTask)
	return pjob, pconn
}
