package maven

import (
	"github.com/samirprakash/boom/pkg/handle"
	"github.com/samirprakash/boom/pkg/task"
)

// Deploy moves your artifact to a remore repository based on the repoId configured in your settings.xml
func Deploy(flags *Flags) {
	if flags.RepoID == "" {
		handle.Info("Missing data - please provide the repository id. \n\nRun `boom maven deploy -h` for usage guidelines!")
		return
	}
	c := "mvn deploy -DrepositoryId=" + flags.RepoID
	task.Execute(c)
}
