package docker

import (
	"fmt"
	"os"

	"github.com/samirprakash/boom/pkg/check"
	"github.com/samirprakash/boom/pkg/task"
	"gopkg.in/src-d/go-git.v4"
)

// SetupContainerEnv stands up a docker container environement and verifies if the containers are ready for use by doing helthchecks
func SetupContainerEnv(flags *Flags) {
	composeFile := flags.ComposeFile
	healthcheckPorts := flags.HealthCheckPorts
	if composeFile == "" {
		fmt.Fprintln(os.Stderr, "\nMissing data - please provide the docker compose file. \nRun `boom docker compose -h` for usage guidelines!")
		return
	} else if healthcheckPorts == "" {
		fmt.Fprintln(os.Stderr, "\nMissing data - please provide the healthcheck ports exposed in the docker compose file. \nRun `boom docker compose -h` for usage guidelines!")
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
		if err != nil {
			fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
			return
		}
	} else {
		fmt.Println("repository that is being cloned already exists on the build environment")
		r, err := git.PlainOpen(path)
		if err != nil {
			fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
		}
		w, err := r.Worktree()
		if err != nil {
			fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
		}
		err = w.Pull(&git.PullOptions{RemoteName: "origin"})
		if err != nil {
			fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
		}
	}

	setupEnvironment := "docker-compose -f " + composeFile + " up --build --detach --remove-orphans"
	task.Execute(setupEnvironment)
	// check if the docker containers are healthy or not based on the ports that have been exposed from docker-compose.yaml
	check.IfDockerComposeResponds(healthcheckPorts)
}
