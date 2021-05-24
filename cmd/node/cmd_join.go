package main

import "fmt"

func joinCmd(secret string, tags []string, hubAddr string) {
	fmt.Printf("%s %v %s", secret, tags, hubAddr)
}
