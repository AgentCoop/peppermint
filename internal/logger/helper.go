package logger

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
	"github.com/fatih/color"
	"time"
)

func RegisterStdoutLogger(key interface{}, prewordColor color.Attribute, utfIcon string, on bool) {
	job.RegisterLogger(key, func(args...interface{}) {
		now := time.Now()
		color := color.New(prewordColor, color.Bold)
		color.Printf("[ %s  %s] → ", utfIcon, now.Format(time.StampMilli))
		fmtStr := args[0].(string)
		if len(args) == 1 {
			fmt.Println(fmtStr)
		} else {
			fmt.Printf(fmtStr + "\n", args[1:]...)
		}
	}, on)
}
