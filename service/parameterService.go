package service

import (
	"bufio"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"

	"github.com/skyhook-cli/skyhook-cli-go/model"
)

var optionsToComponents map[string]func() []model.Prompt = map[string]func() []model.Prompt{
	"platform": func() []model.Prompt {
		return model.InfraContainerPlatform()
	},
	"registry": func() []model.Prompt {
		return model.InfraRegistry()
	},
	"jenkins": func() []model.Prompt {
		return model.InfraJenkins()
	},
	"sonar": func() []model.Prompt {
		return model.InfraSonar()
	},
	"anchore": func() []model.Prompt {
		return model.InfraAnchore()
	},
	"selenium": func() []model.Prompt {
		return model.InfraSelenium()
	},
}

/*
GenerateConfig gets the necessary prompts for the project type and generates the rest of the config struct
*/
func GenerateConfig(config *model.Config, reader *bufio.Reader) {

	prompts := getPrompts(config)

	for i := range prompts {
		p := &prompts[i]
		p.PrintPrompt()
		p.ReadResponse(reader)

		config.Parameters[p.Name] = p.Response
	}

}

func getPrompts(config *model.Config) []model.Prompt {

	switch config.ProjectType {
	case "infra":
		return infraPrompts(config)
	case "app":
		return model.InitializeAppPrompts()
	default:
		return []model.Prompt{}
	}
}

func infraPrompts(config *model.Config) []model.Prompt {
	if len(os.Args) > 2 {

		prompts := []model.Prompt{}

		if _, ok := config.Parameters["projectName"]; !ok {
			prompts = append(prompts, model.InfraProjectName()...)
		}

		for _, v := range os.Args[2:] {
			f, ok := optionsToComponents[v]

			if ok {
				prompts = append(prompts, f()...)
			} else {
				fmt.Println(aurora.Sprintf(aurora.BrightYellow("Warning: unrecognized component %v"), v))
			}
		}

		return prompts
	}
	return model.InitializeInfraPrompts()
}
