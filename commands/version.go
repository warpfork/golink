package commands

import (
	. "fmt"
)

const Version = "1.0.0"

type VersionCmdOpts struct { }

//Version command
func (opts *VersionCmdOpts) Execute(args []string) error {
	Println("golink version", Version)
	return nil
}
