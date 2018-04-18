package docker

import (
	"os"

	"github.com/samirprakash/boom/pkg/task"
	log "github.com/sirupsen/logrus"
)

// ExecuteNewmanTests runs the integration test on the exisiting docker compose environment using newman CLI
func ExecuteNewmanTests(flags *Flags) {
	networkBridge := flags.NetworkBridge
	testCollection := flags.TestCollection
	environmentSpec := flags.EnvironmentSpec

	if networkBridge == "" {
		log.Info("\nMissing data - please provide network bridge name. \nRun `docker network ls` to get a list of existing network bridges!")
		log.Info("\nIf network bridge does not exist then execute `boom docker compose` before running this command!")
		log.Fatal("\nRun `boom docker run -h` for usage guidelines!")
	} else if testCollection == "" {
		log.Fatal("\nMissing data - please provide the JSON file name for your test collection. \nRun `boom docker run -h` for usage guidelines!")
	} else if environmentSpec == "" {
		log.Fatal("\nMissing data - please provide the JSON file defining your execution environment. \nRun `boom docker run -h` for usage guidelines!")
	}

	pwd, _ := os.Getwd()
	v := pwd + "/integration-tests:/etc/postman postman/newman_alpine33:3.8.3"
	c := "/etc/postman/" + testCollection
	e := "/etc/postman/" + environmentSpec
	runTests := "docker run --network " + networkBridge + " -v " + v + " -c=" + c + " -e=" + e
	task.Execute(runTests)
}
