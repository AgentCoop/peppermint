package cmd

import "fmt"

func JoinCmd(secret string, tags []string, hubAddr string) {
	fmt.Printf("%s %v %s", secret, tags, hubAddr)
}
