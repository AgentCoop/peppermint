package utils

import (
	"fmt"
	"os"
)

func FS_FileOrDirExists(pathname string) error {
	if _, err := os.Stat(pathname); os.IsNotExist(err) {
		return fmt.Errorf("fs: pathname '%s' does not exist", pathname)
	}
	return nil
}

