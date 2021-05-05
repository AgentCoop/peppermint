package main

const (
	InstallDir = "/opt/peppermint/cdn"
	DbFilename = "cdn.db"
)

func main() {
	parseCliOptions()
	createDb()
}
