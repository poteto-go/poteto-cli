package cmdrun

import (
	"fmt"
	"os"

	"github.com/poteto-go/poteto-cli/core"
	"github.com/poteto-go/poteto-cli/engine"
	"github.com/poteto-go/poteto/utils"
)

func loadOption() core.RunnerOption {
	configFile, err := os.Open("./poteto.yaml")
	defer configFile.Close()

	if err != nil {
		utils.PotetoPrint("you can use poteto.yaml\n")
		return core.DefaultRunnerOption
	}

	configBytes := make([]byte, 1024)
	n, err := configFile.Read(configBytes)
	if err != nil || n == 0 {
		utils.PotetoPrint("warning error on reading poteto.yaml, use default option\n")
		return core.DefaultRunnerOption
	}

	var option core.RunnerOption
	err = utils.YamlParse(configBytes[:n], &option)
	if err != nil {
		utils.PotetoPrint("warning error on reading poteto.yaml, use default option\n")
		return core.DefaultRunnerOption
	}

	return option
}

func CommandRun() {
	option := loadOption()

	utils.PotetoPrint("You can also use poteto-cli run -h | --help\n")
	for i := 2; i < len(os.Args); i++ {
		switch {
		case os.Args[i] == "-h", os.Args[i] == "--help":
			help()
			os.Exit(-1)
		default:
			utils.PotetoPrint(fmt.Sprintf("unknown command or option: %s\n", os.Args[i]))
			os.Exit(-1)
		}
	}

	engine.RunRun(option)
}

func help() {
	utils.PotetoPrint("poteto-cli run: hot-reload run api server\n")
	utils.PotetoPrint("https://github.com/poteto-go/poteto\n")
	utils.PotetoPrint("=========================================\n")
	utils.PotetoPrint("\n")
	utils.PotetoPrint("Options:\n")
	utils.PotetoPrint("  -h, --help: Display help (this is this)\n")
}
