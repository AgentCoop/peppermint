package logger

import (
	"github.com/AgentCoop/peppermint/internal/logger"
	"github.com/fatih/color"
)

type TestLoggerKey string

var (
	Info = TestLoggerKey("test-info")
)

func init() {
	pic := "🛠"
	logger.RegisterStdoutLogger(Info, color.FgHiYellow, pic, true)
}
