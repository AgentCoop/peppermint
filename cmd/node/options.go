package main

type Version struct {
	Verbose []bool `short:"v" long:"verbose" description:""`
}

type Join struct {
	Tags []string `long:"tag"`
	Hub string `long:"hub" required:"true"`
}

type DbMigrate struct {
	Drop bool `long:"drop" description:"Drop database before migration"`
}

type Run struct {
	HubPort int `long:"hub-port" default:"11000"`
	WebProxyPort int `long:"wp-port" default:"443"`
}

var (
	options = struct {
		DbMigrate `command:"db_migrate"`
		Run `command:"run"`
		Join `command:"join"`
		Version `command:"version"`
	}{}
)
