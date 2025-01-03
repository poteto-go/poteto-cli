package main

import (
	"fmt"
	"os"

	cmdnew "github.com/poteto-go/poteto/cmd/cmd-new"
	cmdrun "github.com/poteto-go/poteto/cmd/cmd-run"
)

func main() {
	if len(os.Args) == 1 {
		help()
		os.Exit(-1)
	}

	if os.Args[1] == "new" {
		cmdnew.CommandNew()
		os.Exit(-1)
	}

	if os.Args[1] == "run" {
		cmdrun.CommandRun()
		os.Exit(-1)
	}

	for i := 1; i < len(os.Args); i++ {
		switch {
		case os.Args[i] == "-h", os.Args[i] == "--help":
			help()
			os.Exit(-1)
		default:
			fmt.Println("unknown command or option:", os.Args[i])
			os.Exit(-1)
		}
	}
}

func help() {
	fmt.Println("poteto-cli: support creating poteto-app")
	fmt.Println("https://github.com/poteto-go/poteto")
	fmt.Println("========================================")
	fmt.Println("")
	fmt.Println("Command: poteto-cli [command]")
	fmt.Println("  run:        hot-reload run golang app")
	fmt.Println("  new:        create new poteto app")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -h, --help: Display help (this is this)")
}
