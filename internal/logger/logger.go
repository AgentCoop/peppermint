package logger

import "github.com/fatih/color"

type loggerKey string

var (
	Debug = loggerKey("debug")
	Error = loggerKey("error")
)

// 🖴 ⚙ 🛠 🛈 ℹ 💻 ⚠ ☠ 🕱
func init() {
	RegisterStdoutLogger(Debug, color.FgHiBlack, "🛠", true)
	RegisterStdoutLogger(Error, color.FgHiRed, "☠", true)
}
