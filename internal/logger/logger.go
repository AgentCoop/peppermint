package logger

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
)

type loggerKey string

var (
	DbKey = loggerKey("db")
)

func init() {
	job.RegisterLogger(DbKey, func(args...interface{}) {
		fmtStr := "[ 🖴 ] → " + args[0].(string)
		if len(args) == 1 {
			fmt.Println(fmtStr)
		} else {
			fmt.Printf(fmtStr, args[1:]...)
		}
	}, true)
}
