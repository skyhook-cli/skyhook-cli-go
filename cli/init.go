package cli

import (
	"bufio"
	"os"

	"github.com/skyhook-cli/skyhook-cli-go/model"
	"github.com/skyhook-cli/skyhook-cli-go/service"
)

/*
RunInit is the entrypoint for creating a new project
*/
func RunInit() {

	reader := bufio.NewReader(os.Stdin)

	config := model.Config{}
	config.ReadConfig()

	projectType := ""

	if len(os.Args) <= 2 {
		infraOrApp := model.InitializeInfraOrApp()
		infraOrApp.PrintPrompt()
		infraOrApp.ReadResponse(reader)

		projectType = infraOrApp.Response
	} else {
		projectType = "infra"
	}

	config.ProjectType = projectType

	service.GenerateConfig(&config, reader)

	config.SaveConfig()
}
