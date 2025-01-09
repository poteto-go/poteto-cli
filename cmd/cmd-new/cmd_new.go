package cmdnew

import (
	"fmt"
	"os"

	"github.com/poteto-go/poteto-cli/engine"
	"github.com/poteto-go/poteto/utils"

	"github.com/manifoldco/promptui"
)

func CommandNew() {
	param := engine.EngineNewParam{}

	utils.PotetoPrint("You can also use poteto-cli new -h | --help\n")
	for i := 2; i < len(os.Args); i++ {
		switch {
		case os.Args[i] == "-h", os.Args[i] == "--help":
			help()
			os.Exit(-1)
		case os.Args[i] == "-f", os.Args[i] == "--fast":
			param.IsFast = true
		case os.Args[i] == "-d", os.Args[i] == "--docker":
			param.IsDocker = true
		case os.Args[i] == "-j", os.Args[i] == "--jsonrpc":
			param.IsJSONRPC = true
		default:
			utils.PotetoPrint(fmt.Sprintf("unknown command or option: %s\n", os.Args[i]))
			os.Exit(-1)
		}
	}

	wd, _ := os.Getwd()
	utils.PotetoPrint(fmt.Sprintf("Generate New Poteto App @%s\n", wd))

	prompt := promptui.Prompt{
		Label: "your project [github.com/github/poteto-api]", // 表示する文言
	}
	projectName, _ := prompt.Run()
	if len(projectName) == 0 {
		projectName = "github.com/github/poteto-api"
	}

	param.ProjectName = projectName

	err := engine.RunNew(param)
	if err != nil {
		panic(err)
	}
}

func help() {
	utils.PotetoPrint("poteto-cli new: support creating poteto-app\n")
	utils.PotetoPrint("https://github.com/poteto-go/poteto\n")
	utils.PotetoPrint("===========================================\n")
	utils.PotetoPrint("\n")
	utils.PotetoPrint("Options:\n")
	utils.PotetoPrint("  -h, --help: Display help (this is this)\n")
	utils.PotetoPrint("  -f, --fast: fast mode api (doesn't gen requestId automatic)\n")
	utils.PotetoPrint("  -d, --docker: with Dockerfile & docker-compose w golang@1.23\n")
	utils.PotetoPrint("  -j, --jsonrpc: jsonrpc template\n")
}
