package maven

import (
	"fmt"
	"os"

	"github.com/samirprakash/boom/pkg/task"
	"github.com/spf13/cobra"
)

// Deploy moves your artifact to a remore repository based on the repoId configured in your settings.xml
func Deploy(cmd *cobra.Command, args []string) {
	repoID := args[0]
	if repoID == "" {
		fmt.Fprintln(os.Stderr, "Missing data - please provide the repository id. \n\nRun `boom maven deploy -h` for usage guidelines!")
		return
	}
	c := "mvn deploy -DrepositoryId=" + repoID
	task.Execute(c)
}
