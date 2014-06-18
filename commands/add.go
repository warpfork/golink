package commands

import (
	. "fmt"
	"path/filepath"
	"strings"

	. "polydawn.net/golink/util"
	. "polydawn.net/pogo/gosh"
)

type AddCmdOpts struct { }

//Runs a container
func (opts *AddCmdOpts) Execute(args []string) error {
	if len(args) != 2 {
		ExitGently("Create takes two arguments: clone URI, and package name.")
	}

	// Load args
	uri := args[0]
	pkg := args[1]
	Println("Adding " + pkg + " from " + uri)

	// URLs are not platform-specific; use strings not filepath
	pkgs := strings.Split(pkg, "/")

	// Array trickery so unrolling works
	srcFolders := []string{ ".", ".gopath", "src" }
	pkgFolders := append(srcFolders, pkgs...)
	pkgFolder := filepath.Join(pkgFolders...)

	// Add submodule to correct location
	Sh("git")(DefaultIO)("submodule", "add", uri, pkgFolder)()

	return nil
}
