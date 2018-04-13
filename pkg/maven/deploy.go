package maven

import (
	"fmt"
	"os"

	"github.com/samirprakash/boom/pkg/task"
)

// Deploy moves your artifact to a remore repository based on the repoId configured in your settings.xml
func Deploy(flags *Flags) {
	if flags.RepoID == "" {
		fmt.Fprintln(os.Stderr, "Missing data - please provide the repository id. \n\nRun `boom maven deploy -h` for usage guidelines!")
		return
	}
	c := "mvn deploy -DrepositoryId=" + flags.RepoID
	task.Execute(c)
}
