package main

var (
	options struct {
		CreateDb struct {
			Force bool `long:"force"`
		} `command:"createdb"`
		JoinCmd struct {
			Secret string `long:"secret"`
		} `command:"join"`
	}
)
