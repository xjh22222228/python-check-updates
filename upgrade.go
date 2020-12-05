package python_check_updates

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/xjh22222228/python-check-updates/constants"
	"log"
	"os/exec"
	"runtime"
)

func isWindows() bool {
	return runtime.GOOS == "windows"
}

func Upgrade()  {
	v := CheckLastVersion()

	if v == constants.Version {
		color.Yellow("Currently the latest version " + v)
		return
	}

	if isWindows() {
		return
	}

	cmd := exec.Command("/bin/bash", "-c", "./install.sh")

	cmd.Stdout = cmd.Stderr
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatalln("Install Fail(StdoutPipe)", err)
		return
	}

	if err := cmd.Start(); err != nil {
		log.Fatalln("Install Fail(Start)", err)
		return
	}

	for {
		b := make([]byte, 1024)
		_, err := stdout.Read(b)

		if err != nil {
			break
		}

		fmt.Print(string(b))
	}

	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}

