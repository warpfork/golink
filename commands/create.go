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
	if len(args) != 1 {
		ExitGently("Init one argument, the package name.  (The executable name will be the last chunk of the package.)")
	}

	// Load args
	pkg  := args[0]
	pkgChunks := strings.Split(pkg, "/")
	name := pkgChunks[len(pkgChunks)-1]
	Println("Creating package " + pkg)

	// URLs are not platform-specific; use strings not filepath
	pkgs := strings.Split(pkg, "/")

	// Array trickery so unrolling works
	srcFolders := []string{ ".", ".gopath", "src" }
	pkgFolders := append(srcFolders, pkgs...)
	pkgFolders = pkgFolders[:len(pkgFolders)-1] // discard the last one because that's where the symlink goes
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
	WriteFile(name + ".go", MainTemplate, 0644)

	// Write goad bash script
	goad := strings.Replace(GoadTemplate, "PACKAGE", pkg,  -1)
	WriteFile(filepath.Join(".", "goad"), goad, 0755)

	// Run build
	Println("Building example...")
	Sh(Abs("goad"))("build")(DefaultIO)()
	Sh(Abs(name))(DefaultIO)()

	return nil
}
