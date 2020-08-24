package cli

import (
	"bufio"
	"os"

	"github.com/skyhook-cli/skyhook-cli-go/model"
)

/*
RunInit is the entrypoint for creating a new project
*/
func RunInit() {
	p := model.Prompt{
		Name:     "InfraOrApp",
		Question: "Is this an Infrastructure repo or an Application repo?",
		Response: "",
		Choices:  []string{"infra", "app"},
		Default:  "infra",
	}
	a := []model.Prompt{p}
	reader := bufio.NewReader(os.Stdin)
	for i := range a {
		p := &a[i]
		p.PrintPrompt()
		p.ReadResponse(reader)
	}

	var c model.Config
	if a[0].Response == "infra" {
		c = model.InfraConfig{
			InfraOrApp: a[0].Response,
		}
	} else if a[0].Response == "app" {
		c = model.AppConfig{
			InfraOrApp: a[0].Response,
		}
	}

	c.SaveConfig()
}
