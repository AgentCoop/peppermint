package main

import (
	"github.com/jessevdk/go-flags"
	"os"
)

var CliOptions struct {
	CreateDb bool `long:"create-db"`
}

func parseCliOptions() {
	parser := flags.NewParser(&CliOptions, flags.PassDoubleDash | flags.PrintErrors | flags.IgnoreUnknown)
	_, err := parser.ParseArgs(os.Args)
	if err != nil { panic(err) }
}
