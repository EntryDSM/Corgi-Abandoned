package init

import (
	"os/exec"
	"runtime"
	"strconv"
)

const (
	windows = "windows"
	root    = 0
)

func Init() {
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

func isUnixSuperuser() bool {
	const errorMessage = "An error occured while checking application's permission."

	cmd := exec.Command("id", "-u")
	output, err := cmd.Output()

	if err != nil {
		panic(errorMessage)
	}

	i, err := strconv.Atoi(string(output[:len(output)-1]))

	if err != nil {
		panic(errorMessage)
	}

	return i == root
}

func isWindows() bool {
	return runtime.GOOS == windows
}
