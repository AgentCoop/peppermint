package logger

import "github.com/fatih/color"

type loggerKey string

var (
	Debug = loggerKey("debug")
)

// 🖴 ⚙ 🛠 🛈 ℹ 💻
func init() {
	RegisterStdoutLogger(Debug, color.FgHiBlack, "🛠", true)
}
