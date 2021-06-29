package test

import (
	cli "github.com/AgentCoop/peppermint/internal/app/cli_client"
)

const (
	CMD_NAME_SINGLE           = "single"
	CMD_NAME_SINGLE_ENCRYPTED = "single-encrypted"
	CMD_NAME_STREAMABLE       = "streamable"
)

type Single struct{}
type SingleEncrypted struct{}
type Streamable struct{}

var (
	options = struct {
		cli.CommonOptions
		Timeout           uint   `long:"timeout" short:"t"`
		RsDelay           uint32 `long:"rs-delay"`
		RsDelayJitter     uint32 `long:"rs-delay-jitter"`
		Token             string `long:"token"` // A random word
		RsBulkDataMin     uint32 `long:"rs-bulk-min"`
		RsBulkDataMax     uint32 `long:"rs-bulk-max"`
		RqBulkDataMin     uint32 `long:"rq-bulk-min"`
		RqBulkDataMax     uint32 `long:"rq-bulk-max"`
		CallFailureProbab uint   `long:"call-failure-prob"` // 0-100
		CallRepeatCount   uint   `long:"repeat" short:"r"`  // Specifies how many times a call must be invoked
		CallWorkersCount  uint   `long:"workers" short:"w"` // Specifies the number of goroutines executing the call
		Single            `command:"single"`
		SingleEncrypted   `command:"single-encrypted"`
		Streamable        `command:"streamable"`
	}{}
)
