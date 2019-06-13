package initialize

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/EntryDSM/Corgi/config"
)

const (
	windows = "windows"
	root    = 0
)

func Initialize() {
	checkUserPrivilege()
	initializeConfig()
}

func checkUserPrivilege() {
	isSuperuser := isApplicationRunningOnSuperuser()
	if !isSuperuser {
		panic("Corgi needs root privilege.")
	}
}

func initializeConfig() {
	configFilePath := os.Args[0]
	config.InitializeConfig(configFilePath)
}

func isApplicationRunningOnSuperuser() bool {
	if isWindows() {
		panic("Corgi doesn't support windows platform.")
	} else {
		return isUnixSuperuser()
	}
}

func isUnixSuperuser() bool {
	const errorMessage = "An error occured while checking application's permission."

	cmd := exec.Command("id", "-u")
	rawUIDString, err := cmd.Output()

	if err != nil {
		panic(errorMessage)
	}

	uidString := strings.TrimSpace(string(rawUIDString))
	uid, err := strconv.Atoi(uidString)

	if err != nil {
		panic(errorMessage)
	}

	return uid == root
}

func isWindows() bool {
	return runtime.GOOS == windows
}
