package logger

import (
	"github.com/AgentCoop/peppermint/internal/logger"
	"github.com/fatih/color"
)

type HubLoggerKey string

var (
	InfoKey = HubLoggerKey("hub-info")
	WarnKey = HubLoggerKey("hub-warn")
)

func init() {
	pic := "🖧"
	logger.RegisterStdoutLogger(InfoKey, color.FgGreen, pic, true)
	logger.RegisterStdoutLogger(WarnKey, color.FgHiRed, pic, true)
}
