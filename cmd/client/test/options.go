package main

import "github.com/AgentCoop/peppermint/internal/service/test/grpc/client"

var (
	options = struct {
		Count                 int    `long:"count" short:"c"`
		Service               string `long:"service" short:"s"`
		client.CmdPingOptions `command:"ping"`
	}{}
)
