package maven

import (
	"github.com/samirprakash/boom/pkg/check"
	"github.com/samirprakash/boom/pkg/task"
	log "github.com/sirupsen/logrus"
)

// Deploy moves your artifact to a remore repository based on the repoId configured in your settings.xml
func Deploy(flags *Flags) {
	if flags.RepoID == "" {
		log.Fatal("Missing data - please provide the repository id. \n\nRun `boom maven deploy -h` for usage guidelines!")
	}
	if check.IsBranchMorD() {
		c := "mvn deploy -DrepositoryId=" + flags.RepoID
		task.Execute(c)
	}
}
