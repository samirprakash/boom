package task

import (
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Execute is a generic implementation to execute the commands that are provided
// it tokenizes the incoming string command to generate executable commands
// and spits the output to the console
func Execute(c string) {
	tokens := strings.Fields(c)
	executable := tokens[0]
	args := tokens[1:len(tokens)]

	cmd := exec.Command(executable, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
