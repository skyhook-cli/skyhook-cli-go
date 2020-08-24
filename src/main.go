package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/logrusorgru/aurora"
	. "github.com/skyhook-cli/skyhook-cli-go/model"
	. "github.com/skyhook-cli/skyhook-cli-go/util"
)

func main() {

	Clear()
	fmt.Println(Bold(Cyan("=========================================================")))
	fmt.Println(Bold(Cyan("#                      Skyhook CLI                      #")))
	fmt.Println(Bold(Cyan("=========================================================")))

	var command string
	if len(os.Args) > 1 {
		command = os.Args[1]
	} else {
		command = "init"
	}

	switch command {
	case "init":
		p := Prompt{
			Name:     "InfraOrApp",
			Question: "Is this an Infrastructure repo or an Application repo?",
			Response: "",
			Choices:  []string{"infra", "app"},
			Default:  "infra",
		}
		a := []Prompt{p}
		reader := bufio.NewReader(os.Stdin)
		for i, _ := range a {
			p := &a[i]
			p.PrintPrompt()
			p.ReadResponse(reader)
		}

		var c Config
		if a[0].Response == "infra" {
			c = InfraConfig{
				InfraOrApp: a[0].Response,
			}
		} else if a[0].Response == "app" {
			c = AppConfig{
				InfraOrApp: a[0].Response,
			}
		}

		c.SaveConfig()
	default:
		fmt.Println(Sprintf(Bold(Red("unknown command '%v'")), command))
		fmt.Println(Yellow("only 'init' is implemented so far"))
	}
}
