package main

type Version struct {
	Verbose []bool `short:"v" long:"verbose" description:""`
}

var (
	options = struct {
		DbMigrate struct {
			Force bool `long:"force"`
		} `command:"db_migrate"`
		Run struct {
			HubPort int `long:"hub-port" default:"11000"`
			WebProxyPort int `long:"wp-port" default:"443"`
		} `command:"run"`
		Join struct {
			Secret string `long:"secret"`
		} `command:"join"`
		Version `command:"version"`
	}{}
)
