package app

import (
	"fmt"
	"os"
)

func ProfileFromEnv() AppProfile {
	var profile AppProfile
	env := os.Getenv(ENV)
	switch {
	case len(env) == 0:
		profile = AppProfiles[PROD]
	default:
		p, ok := AppProfiles[EnvString(env)]
		if !ok {
			panic(fmt.Errorf("invalid PEPPERMINT environment variable value %s", env))
		}
		profile = p
	}
	return profile
}
