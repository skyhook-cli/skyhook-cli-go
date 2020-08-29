package model

/*
InitializeAppPrompts creates the prompts for app params
*/
func InitializeAppPrompts() []Prompt {
	return []Prompt{
		{
			Name:     "containerPlatform",
			Question: "Which container management platform are you using?",
			Choices:  []string{"EKS", "ECS Fargate", "ECS EC2", "OCP"},
			Default:  "EKS",
		},
		{
			Name:     "dockerImage",
			Question: "Enter the docker image name for this app (e.g. registry/namespace/image:tag):",
		},
		{
			Name:     "appName",
			Question: "Enter the desired app name - will be used for service/task name:",
		},
	}
}
