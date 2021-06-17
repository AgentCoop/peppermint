package logger

import (
	"github.com/AgentCoop/peppermint/internal/logger"
	"github.com/fatih/color"
)

type HubLoggerKey string

var (
	Info = HubLoggerKey("hub-info")
	Warn = HubLoggerKey("hub-warn")
)

func init() {
	pic := "🖧"
	logger.RegisterStdoutLogger(Info, color.FgGreen, pic, true)
	logger.RegisterStdoutLogger(Warn, color.FgHiRed, pic, true)
}
