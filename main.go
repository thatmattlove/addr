package main

import (
	_ "embed"

	"github.com/thatmattlove/addr/cmd"
)

//go:embed .version
var Version string

func main() {
	cmd.Init(Version).Execute()
}
