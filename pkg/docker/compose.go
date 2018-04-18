package docker

import (
	"os"

	"github.com/samirprakash/boom/pkg/check"
	"github.com/samirprakash/boom/pkg/task"
	log "github.com/sirupsen/logrus"
)

// SetupContainerEnv stands up a docker container environement and verifies if the containers are ready for use by doing helthchecks
func SetupContainerEnv(flags *Flags) {
	composeFile := flags.ComposeFile
	healthcheckPorts := flags.HealthCheckPorts
	repoName := flags.RepoName
	if composeFile == "" {
		log.Fatal("\nMissing data - please provide the docker compose file. \nRun `boom docker compose -h` for usage guidelines!")
	} else if healthcheckPorts == "" {
		log.Fatal("\nMissing data - please provide the healthcheck ports exposed in the docker compose file. \nRun `boom docker compose -h` for usage guidelines!")
	}

	// clone config source repo if not already present in the build environment
	path := os.Getenv("TC_CONFIG_PATH")
	repo, _ := check.IfDirExists(path)
	if !repo {
		if repoName == "" {
			log.Fatal("\nMissing data - please provide the repo name to be cloned. \nRun `boom docker compose -h` for usage guidelines!")
		}
		task.Clone(path, repoName)
	}
	task.Fetch(path)

	setupEnvironment := "docker-compose -f " + composeFile + " up --build --detach --remove-orphans"
	task.Execute(setupEnvironment)
	// check if the docker containers are healthy or not based on the exposed ports
	check.IfDockerComposeResponds(healthcheckPorts)
}
