package main

import "github.com/samirprakash/go-boom/cmd"

var (
	// VERSION would be provided at build
	VERSION = "0.0.1"
)

func main() {
	cmd.Execute(VERSION)
}
