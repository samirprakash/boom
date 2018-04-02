package maven

import (
	"github.com/samirprakash/boom/pkg/task"
	"github.com/spf13/cobra"
)

// Clean removes existing target folder from your maven based code base
func Clean(cmd *cobra.Command, args []string) {
	c := "mvn clean"
	task.Execute(c)
}
