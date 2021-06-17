package utils

import "strings"

func Str_SplitAndLast(name string, sep string) string {
	parts := strings.Split(name, sep)
	return parts[len(parts)-1]
}
