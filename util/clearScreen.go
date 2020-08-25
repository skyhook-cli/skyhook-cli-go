package util

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/logrusorgru/aurora"
)

/*
Clear the terminal
*/
func Clear() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = nil
	}

	if cmd != nil {
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Println(aurora.Yellow("Warning:"), "cannot clear console")
	}
}
