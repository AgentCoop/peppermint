package main

import (
	"github.com/AgentCoop/peppermint/internal/app/cli_client/test"
)

func main() {
	appTest := test.NewTestApp()
	j := appTest.Job()
	// Execute command
	<-j.Run()
	// Handle error
	_, jobErr := j.GetInterruptedBy()
	if jobErr != nil {
		panic(jobErr)
	}
}
