package str

import "strings"

func SplitAndLast(name string, sep string) string {
	parts := strings.Split(name, sep)
	return parts[len(parts)-1]
}

type SplitPartPos int

const (
	Begin SplitPartPos = iota
	End
)

func Split(name string, sep string, ) string {
	parts := strings.Split(name, sep)
	return parts[len(parts)-1]
}

