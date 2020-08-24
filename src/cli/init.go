package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"

	"github.com/skyhook-cli/skyhook-cli-go/model"
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
		InfraOrApp: infraOrApp.Response,
		Parameters: make(map[string]string),
	}
	var params []model.Prompt

	switch infraOrApp.Response {
	case "infra":
		params = model.InitializeInfraPrompts()
	case "app":
		params = model.InitializeAppPrompts()
	default:
		fmt.Println(aurora.Red("invalid option!"))
		os.Exit(1)
	}

	for i := range params {
		p := &params[i]
		p.PrintPrompt()
		p.ReadResponse(reader)

		config.Parameters[p.Name] = p.Response
	}

	config.SaveConfig()
}
