package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func IntToHex(int interface{}, width int) string {
	return fmt.Sprintf(fmt.Sprintf("%%.%dx", width), int)
}

func Hex2int(hexStr string) uint64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, err := strconv.ParseUint(cleaned, 16, 64)
	if err != nil {
		panic("utils.Hex2int conversion failed")
	}

	return uint64(result)
}
