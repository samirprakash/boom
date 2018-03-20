package cmd

import (
	"os/exec"

	"github.com/samirprakash/go-boom/helper"
	"github.com/spf13/cobra"
)

// verifyCmd represents the "mvn verify" command
// If executed from the root location of pom.xml,
// it would verify the tests and generate an output report
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify your test case and generate a report out of it",
	Run: func(cmd *cobra.Command, args []string) {
		verifyTests()
	},
}

func verifyTests() {
	msg := "Verifying your tests ... "

	// Spinner with custom message to display execution progress
	s := helper.StartSpinner(msg)

	cmd := exec.Command("mvn", "verify")
	helper.PrintCommand(cmd)

	output, err := cmd.CombinedOutput()
	if err != nil {
		helper.PrintError(err)
	}

	s.Stop()
	helper.PrintOutput(output)
}

func init() {
	mavenCmd.AddCommand(verifyCmd)
}
