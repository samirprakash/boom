package maven

import (
	"github.com/samirprakash/boom/pkg/task"
	"github.com/spf13/cobra"
)

// Verify defines the option to validate and compile maven based code base
func Verify(cmd *cobra.Command, args []string) {
	c := "mvn verify"
	task.Execute(c)
}
