package maven

import (
	"github.com/samirprakash/boom/utils"
	"github.com/spf13/cobra"
)

func Verify(cmd *cobra.Command, args []string) {
	c := "mvn verify"
	utils.Execute(c)
}
