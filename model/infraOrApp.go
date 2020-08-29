package model

/*
InitializeInfraOrApp creates the first prompt
*/
func InitializeInfraOrApp() Prompt {
	return Prompt{
		Name:     "ProjectType",
		Question: "Is this an Infrastructure project or an Application project?",
		Choices:  []string{"infra", "app"},
		Default:  "infra",
	}
}
