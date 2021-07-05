package fs

import (
	"os"
)

func IfExists(pathname string) bool {
	if _, err := os.Stat(pathname); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
