package commands

import (
	. "fmt"
	"path/filepath"
	"strings"

	. "polydawn.net/golink/util"
	. "polydawn.net/pogo/gosh"
)

type CreateCmdOpts struct { }

//Runs a container
func (opts *CreateCmdOpts) Execute(args []string) error {
	if len(args) != 2 {
		ExitGently("Init requires two arguments: executable name, and package name.")
	}

	// Load args
	name := args[0]
	pkg  := args[1]
	Println("Creating " + name + " in package " + pkg)

	// URLs are not platform-specific; use strings not filepath
	pkgs := strings.Split(pkg, "/")

	// Array trickery so unrolling works
	srcFolders := []string{ ".", ".gopath", "src" }
	pkgFolders := append(srcFolders, pkgs...)
	pkgFolder := filepath.Join(pkgFolders...)

	// How many folders up does the self-ref symlink need to go?
	n := len(pkgFolders) - 1
	linkDest := ""
	for ; n > 0 ; n-- {
		linkDest = filepath.Join(linkDest, "..")
	}

	// Create gopath folders with self-ref symlink
	CreateFolder(pkgFolder)
	Symlink(linkDest, filepath.Join(pkgFolder, name))

	// Write main code
	CreateFolder("main")
	WriteFile(filepath.Join("main", name + ".go"), MainTemplate, 0644)

	// Write goad bash script
	goad := strings.Replace(GoadTemplate, "PACKAGE", pkg,  -1)
	goad  = strings.Replace(goad,         "NAME",    name, -1)
	WriteFile(filepath.Join(".", "goad"), goad, 0755)

	// Run build
	Println("Building example...")
	Sh(Abs("goad"))(DefaultIO)()
	Sh(Abs(name))(DefaultIO)()

	return nil
}
