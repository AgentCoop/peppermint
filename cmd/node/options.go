package main

var (
	options = struct {
		CreateDb struct {
			Force bool `long:"force"`
		} `command:"create_db"`
		Join struct {
			Secret string `long:"secret"`
		} `command:"join"`
	}{}
)
