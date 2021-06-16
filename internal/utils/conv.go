package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrConv_HexToInt = errors.New("utils.Conv_HexToInt: failed")
)

func Conv_IntToHex(int interface{}, width int) string {
	return fmt.Sprintf(fmt.Sprintf("%%.%dx", width), int)
}

func Conv_HexToInt(hexStr string) uint64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)
	// base 16 for hexadecimal
	result, err := strconv.ParseUint(cleaned, 16, 64)
	if err != nil { panic(ErrConv_HexToInt) }
	return uint64(result)
}

func Conv_FromLongToShortMethod(name string) string {
	if name[0] != '/' {
		return name
	}
	parts := strings.Split(name, "/")
	return parts[len(parts)-1]
}
