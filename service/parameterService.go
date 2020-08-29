package service

import (
	"bufio"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"

	"github.com/skyhook-cli/skyhook-cli-go/model"
)

/*
GenerateConfig gets the necessary prompts for the project type and generates the rest of the config struct
*/
func GenerateConfig(config *model.Config, reader *bufio.Reader, projectType string) {

	prompts := getPrompts(projectType)

	for i := range prompts {
		p := &prompts[i]
		p.PrintPrompt()
		p.ReadResponse(reader)

		config.Parameters[p.Name] = p.Response
	}

}

func getPrompts(projectType string) []model.Prompt {

	switch projectType {
	case "infra":
		if len(os.Args) > 2 {
			fmt.Println(aurora.Yellow("TODO: implement this scenario"))
			return []model.Prompt{}
		}
		return model.InitializeInfraPrompts()
	case "app":
		return model.InitializeAppPrompts()
	default:
		return []model.Prompt{}
	}
}
