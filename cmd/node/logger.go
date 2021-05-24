package main

import (
	"fmt"
	"github.com/AgentCoop/go-work"
	"strings"
)

func initLogger(logLevel int, ) {
	job.DefaultLogLevel = logLevel
	job.RegisterDefaultLogger(func() job.LogLevelMap {
		m := make(job.LogLevelMap, 3)
		handler := func(record interface{}, level int) {
			prefix := fmt.Sprintf(" 💻[ %s ] ", "") +
				strings.Repeat("☞ ", level)
			fmt.Printf("%s%s\n", prefix, record.(string))
		}
		m[0] = job.NewLogLevelMapItem(make(chan interface{}), handler)
		m[1] = job.NewLogLevelMapItem(make(chan interface{}), handler)
		m[2] = job.NewLogLevelMapItem(make(chan interface{}), handler)
		return m
	})
}
