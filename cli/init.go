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
	infraOrApp := model.InitializeInfraOrApp()

	reader := bufio.NewReader(os.Stdin)
	infraOrApp.PrintPrompt()
	infraOrApp.ReadResponse(reader)

	config := model.Config{
		ProjectType: infraOrApp.Response,
		Parameters:  make(map[string]string),
	}

	service.GenerateConfig(&config, reader, infraOrApp.Response)

	config.SaveConfig()
}
