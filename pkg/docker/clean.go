package docker

import (
	"github.com/samirprakash/boom/pkg/task"
)

// Clean defines the sub-command to remove existing containers which are no more in use
// to provide a clean build environment for re-runs
func Clean() {
	c := "docker-compose stop"
	task.Execute(c)
	c = "docker ps -a -q -f status=exited"
	task.Execute(c)
}
