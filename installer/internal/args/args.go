package args

import (
	"os"
	"path"
	"strings"
)

var config ArgsConfig

type ArgsConfig struct {
	ConfigPath string
}

func init() {
	config = ArgsConfig{}

	args := os.Args[1:]
	for i := 0; i < len(args)-1; i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "--") || strings.HasPrefix(arg, "-") {
			switch arg {
			case "--config":
				config.ConfigPath = parseConfigPath(args[i+1])
			case "-c":
				config.ConfigPath = parseConfigPath(args[i+1])
			}
		}
	}
}

func GetArgs() ArgsConfig {
	return config
}

func parseConfigPath(arg string) string {
	if strings.HasPrefix(arg, "./") {
		cwd, _ := os.Getwd()
		return path.Join(cwd, arg)
	}

	return arg
}
