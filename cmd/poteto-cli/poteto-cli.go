package main

import (
	"fmt"
	"os"

	cmdnew "github.com/poteto-go/poteto-cli/cmd/cmd-new"
	cmdrun "github.com/poteto-go/poteto-cli/cmd/cmd-run"
	"github.com/poteto-go/poteto/constant"
	"github.com/poteto-go/poteto/utils"
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
		case os.Args[i] == "-v", os.Args[i] == "--version":
			utils.PotetoPrint(fmt.Sprintf("poteto-cli version: %s\n", constant.VERSION))
			os.Exit(-1)
		default:
			utils.PotetoPrint(fmt.Sprintf("unknown command or option: %s\n", os.Args[i]))
			os.Exit(-1)
		}
	}
}

func help() {
	utils.PotetoPrint("poteto-cli: support creating poteto-app\n")
	utils.PotetoPrint("https://github.com/poteto-go/poteto\n")
	utils.PotetoPrint("==============================================\n")
	utils.PotetoPrint("\n")
	utils.PotetoPrint("Command: poteto-cli [command]\n")
	utils.PotetoPrint("  run:        hot-reload run golang app\n")
	utils.PotetoPrint("  new:        create new poteto app\n")
	utils.PotetoPrint("\n")
	utils.PotetoPrint("Options:\n")
	utils.PotetoPrint("  -h, --help   : Display help (this is this)\n")
	utils.PotetoPrint("  -v, --version: Display version of poteto-cli\n")
}
