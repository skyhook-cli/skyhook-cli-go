package main

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/skyhook-cli/skyhook-cli-go/cli"
	"github.com/skyhook-cli/skyhook-cli-go/util"
)

func main() {

	util.Clear()
	fmt.Println(aurora.Bold(aurora.Cyan("=========================================================")))
	fmt.Println(aurora.Bold(aurora.Cyan("#                      Skyhook CLI                      #")))
	fmt.Println(aurora.Bold(aurora.Cyan("=========================================================")))

	var command string
	if len(os.Args) > 1 {
		command = os.Args[1]
	} else {
		command = "init"
	}

	switch command {
	case "init":
		cli.RunInit()
	default:
		fmt.Println(aurora.Sprintf(aurora.Bold(aurora.Red("unknown command '%v'")), command))
		fmt.Println(aurora.Yellow("only 'init' is implemented so far"))
	}
}
