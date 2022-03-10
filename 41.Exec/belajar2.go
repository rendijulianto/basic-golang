package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {

	if runtime.GOOS == "windows" {
		output, _ = exec.Command("cmd", "/C", "git config user.name").Output()
	} else {
		output, _ = exec.Command("bash", "-c", "git config user.name").Output()
	}

	fmt.Println(output)
}
