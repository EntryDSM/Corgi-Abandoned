package initialize

import (
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/EntryDSM/Corgi/fetch"

	"github.com/EntryDSM/Corgi/args"
	"github.com/EntryDSM/Corgi/config"
)

const (
	windows = "windows"
	rootUID = 0
)

func Initialize() {
	checkUserPrivilege()
	initializeArguments()
	initializeConfig()
	initializeHTTPClient()
	initializeAuthenticatedRequest()
}

func checkUserPrivilege() {
	isSuperuser := isApplicationRunningOnSuperuser()
	if !isSuperuser {
		panic("Corgi needs root privilege.")
	}
}

func isApplicationRunningOnSuperuser() bool {
	if isWindows() {
		panic("Corgi doesn't support windows platform.")
	} else {
		return isUnixSuperuser()
	}
}

func isWindows() bool {
	return runtime.GOOS == windows
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

	return uid == rootUID
}

func initializeArguments() {
	args.InitializeArguments()
}

func initializeConfig() {
	config.InitializeConfig(*args.ConfigFilePath)
}

func initializeHTTPClient() {
	fetch.InitializeHttpClient()
}

func initializeAuthenticatedRequest() {
	fetch.InitializeRequestWithAuth()
}
