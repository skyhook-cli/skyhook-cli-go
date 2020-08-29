package model

/*
InfraProjectName returns a prompt to get the project name
*/
func InfraProjectName() []Prompt {
	return []Prompt{
		{
			Name:     "projectName",
			Question: "Enter the name of your project (will be prefixed to most resources):",
		},
	}
}

/*
InfraContainerPlatform returns a prompt to get the desired container management platform
*/
func InfraContainerPlatform() []Prompt {
	return []Prompt{
		{
			Name:     "containerPlatform",
			Question: "Which container management platform do you want to use?",
			Choices:  []string{"EKS", "ECS Fargate", "ECS EC2", "OCP"},
			Default:  "EKS",
		},
	}
}

/*
InfraRegistry returns a prompt to get the desired docker container registry
*/
func InfraRegistry() []Prompt {
	return []Prompt{
		{
			Name:     "imageRegistry",
			Question: "Which docker image registry do you want to use?",
			Choices:  []string{"ECR", "Nexus"},
			Default:  "ECR",
		},
	}
}

/*
InfraJenkins returns prompts to get the admin username and password for the jenkins instance
*/
func InfraJenkins() []Prompt {
	return []Prompt{
		{
			Name:     "jenkinsAdminUser",
			Question: "Enter the Jenkins admin username:",
		},
		{
			Name:     "jenkinsAdminPassword",
			Question: "Enter the Jenkins admin password:",
		},
	}
}

/*
InfraSonar returns prompts to get the admin username and password for the sonar instance
*/
func InfraSonar() []Prompt {
	return []Prompt{
		{
			Name:     "sonarAdminUser",
			Question: "Enter the Sonarqube admin username:",
		},
		{
			Name:     "sonarAdminPassword",
			Question: "Enter the Sonarqube admin password:",
		},
	}
}

/*
InfraAnchore returns prompts to get the admin username and password for the anchore instance
*/
func InfraAnchore() []Prompt {
	return []Prompt{
		{
			Name:     "anchoreAdminUsername",
			Question: "Enter the Anchore admin username:",
		},
		{
			Name:     "anchoreAdminPassword",
			Question: "Enter the Anchore admin password:",
		},
	}
}

/*
InfraSelenium returns prompts to get the admin username and password for the selenium instance
*/
func InfraSelenium() []Prompt {
	return []Prompt{
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

/*
InitializeInfraPrompts creates the prompts for infra params
*/
func InitializeInfraPrompts() []Prompt {
	allPrompts := []Prompt{}
	allPrompts = append(allPrompts, InfraProjectName()...)
	allPrompts = append(allPrompts, InfraContainerPlatform()...)
	allPrompts = append(allPrompts, InfraRegistry()...)
	allPrompts = append(allPrompts, InfraJenkins()...)
	allPrompts = append(allPrompts, InfraSonar()...)
	allPrompts = append(allPrompts, InfraAnchore()...)
	allPrompts = append(allPrompts, InfraSelenium()...)
	return allPrompts
}
