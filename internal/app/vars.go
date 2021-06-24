package app

var (
	AppProfiles = map[EnvString]AppProfile{
		DEV: {
			DbFilename: "peppermint-dev.db",
		},
		PROD: {
			DbFilename: "peppermint.db",
		},
		TEST: {
			DbFilename: "peppermint-test.db",
		},
	}
)
