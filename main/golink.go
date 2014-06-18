package main

import (
	. "fmt"
	"os"

	"github.com/jessevdk/go-flags"

	. "polydawn.net/golink/commands"
	. "polydawn.net/golink/util"
)

var parser = flags.NewNamedParser("golink", flags.Default | flags.HelpFlag)

const EXIT_BADARGS = 1
const EXIT_PANIC = 2
const EXIT_BAD_USER = 10

// print only the error message (don't dump stacks).
// unless any debug mode is on; then don't recover, because we want to dump stacks.
func panicHandler() {
	if err := recover(); err != nil {

		//GoLinkError is used for user-friendly exits. Just print & exit.
		if glError, ok := err.(GoLinkError) ; ok {
			Print(glError.Error())
			os.Exit(EXIT_BAD_USER)
		}

		//Check for existence of debug environment variable
		if len(os.Getenv("DEBUG")) == 0 && len(os.Getenv("DEBUG_STACK")) == 0  {
			//Debug not set, be friendlier about the problem
			Println(err)
			Println("\n" + "GoLink crashed!" + "\n" + "To see more about what went wrong, turn on stack traces by running:" + "\n\n" + "export DEBUG=1" + "\n\n" + "Feel free to contact the developers for help:" + "\n" + "https://github.com/polydawn/golink" + "\n")
			os.Exit(EXIT_PANIC)
		} else {
			//Adds main to the top of the stack, but keeps original information.
			//Nothing we can do about it. Golaaaaannngggg....
			panic(err)
		}
	}
}

func main() {
	defer panicHandler()

	// parser.AddCommand(
	// 	"command",
	// 	"description",
	// 	"long description",
	// 	&whateverCmd{}
	// )

	parser.AddCommand(
		"create",
		"Create a new Golang project",
		"Create a new Golang project",
		&CreateCmdOpts{},
	)
	parser.AddCommand(
		"add",
		"Add a new library to a project",
		"Add a new library to a project",
		&AddCmdOpts{},
	)
	parser.AddCommand(
		"version",
		"Print GoLink version",
		"Print GoLink version",
		&VersionCmdOpts{},
	)

	//Go-flags is a little too clever with sub-commands.
	//Check for 'help' manually before args parse
	if len(os.Args) < 2 || os.Args[1] == "help" {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	//Parse for command & flags, and exit with a relevant return code.
	_, err := parser.Parse()
	if err != nil {
		os.Exit(EXIT_BADARGS)
	} else {
		os.Exit(0)
	}
}
