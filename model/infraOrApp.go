package model

/*
InitializeInfraOrApp creates the first prompt
*/
func InitializeInfraOrApp() Prompt {
	return Prompt{
		Name:     "InfraOrApp",
		Question: "Is this an Infrastructure repo or an Application repo?",
		Choices:  []string{"infra", "app"},
		Default:  "infra",
	}
}
