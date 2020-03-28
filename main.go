package main

import "fmt"

var (
	version = "unknown version"
	commit  = "unknown commit"
	date    = "unknown date"
)

func main() {
	fmt.Printf("version: %v, commit: %v, date: %v", version, commit, date)
}
