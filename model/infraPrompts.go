package model

/*
InitializeInfraPrompts creates the prompts for infra params
*/
func InitializeInfraPrompts() []Prompt {
	return []Prompt{
		{
			Name:     "projectName",
			Question: "Enter the name of your project (will be prefixed to most resources):",
		},
		{
			Name:     "containerPlatform",
			Question: "Which container management platform do you want to use?",
			Choices:  []string{"EKS", "ECS Fargate", "ECS EC2", "OCP"},
			Default:  "EKS",
		},
		{
			Name:     "imageRepository",
			Question: "Which docker image repository do you want to use?",
			Choices:  []string{"ECR", "Nexus"},
			Default:  "ECR",
		},
		{
			Name:     "jenkinsAdminUser",
			Question: "Enter the Jenkins admin username:",
		},
		{
			Name:     "jenkinsAdminPassword",
			Question: "Enter the Jenkins admin password:",
		},
		{
			Name:     "sonarAdminUser",
			Question: "Enter the Sonarqube admin username:",
		},
		{
			Name:     "sonarAdminPassword",
			Question: "Enter the Sonarqube admin password:",
		},
		{
			Name:     "anchoreAdminUsername",
			Question: "Enter the Anchore admin username:",
		},
		{
			Name:     "anchoreAdminPassword",
			Question: "Enter the Anchore admin password:",
		},
		{
			Name:     "seleniumAdminUsername",
			Question: "Enter the Selenium admin username:",
		},
		{
			Name:     "seleniumAdminPassword",
			Question: "Enter the Selenium admin password:",
		},
	}
}
