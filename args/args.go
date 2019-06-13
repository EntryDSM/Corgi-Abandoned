package args

import (
	"flag"
	"os"
)

var ConfigFilePath *string

func InitializeArguments() {
	ConfigFilePath = flag.String("config", "", "Config file path.")
	flag.Parse()

	if hasNoArguments() {
		flag.Usage()
		os.Exit(0)
	}
}

func hasNoArguments() bool {
	return flag.NFlag() == 0
}
