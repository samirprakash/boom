package maven

import (
	"github.com/samirprakash/boom/utils"
	"github.com/spf13/cobra"
)

// Validate compiles and validates your maveb based code base
func Validate(cmd *cobra.Command, args []string) {
	c := "mvn validate compile"
	utils.Execute(c)
}
