package main

import "github.com/kawana77b/ghq-ex/cmd"

var Version string = "0.0.1"

func main() {
	cmd.Version = Version
	cmd.Execute()
}
