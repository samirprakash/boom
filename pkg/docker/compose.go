package docker

import (
	"fmt"
	"os"

	"github.com/samirprakash/boom/pkg/check"
	"github.com/samirprakash/boom/pkg/handle"
	"github.com/samirprakash/boom/pkg/task"
	"gopkg.in/src-d/go-git.v4"
)

// SetupContainerEnv stands up a docker container environement and verifies if the containers are ready for use by doing helthchecks
func SetupContainerEnv(flags *Flags) {
	composeFile := flags.ComposeFile
	healthcheckPorts := flags.HealthCheckPorts
	if composeFile == "" {
		handle.Info("\nMissing data - please provide the docker compose file. \nRun `boom docker compose -h` for usage guidelines!")
		return
	} else if healthcheckPorts == "" {
		handle.Info("\nMissing data - please provide the healthcheck ports exposed in the docker compose file. \nRun `boom docker compose -h` for usage guidelines!")
		return
	}

	// clone config source repo if not already present in the build environment
	path := os.Getenv("TC_CONFIG_PATH")
	uname := os.Getenv("GIT_USERNAME")
	pwd := os.Getenv("GIT_PASSWORD")
	repo, _ := check.IfDirExists(path)
	url := "https://" + uname + ":" + pwd + "@github.com/" + flags.CloneURL + ".git"

	if !repo {
		_, err := git.PlainClone(path, false, &git.CloneOptions{
			URL:               url,
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		})
		handle.Error(err)
	} else {
		fmt.Println("repository that is being cloned already exists on the build environment")
		r, err := git.PlainOpen(path)
		handle.Error(err)
		w, err := r.Worktree()
		handle.Error(err)
		err = w.Pull(&git.PullOptions{RemoteName: "origin"})
		if err != nil {
			handle.Warning(err.Error())
		}
	}

	setupEnvironment := "docker-compose -f " + composeFile + " up --build --detach --remove-orphans"
	task.Execute(setupEnvironment)
	// check if the docker containers are healthy or not based on the ports that have been exposed from docker-compose.yaml
	check.IfDockerComposeResponds(healthcheckPorts)
}
