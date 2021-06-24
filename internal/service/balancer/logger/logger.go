package logger

import (
	"github.com/AgentCoop/peppermint/internal/logger"
	"github.com/fatih/color"
)

type HubLoggerKey string

var (
	Info = HubLoggerKey("balancer-info")
)

func init() {
	pic := "🌎"
	logger.RegisterStdoutLogger(Info, color.FgBlue, pic, true)
}
