package cmd

import (
	"os/exec"

	"github.com/samirprakash/go-boom/helper"
	"github.com/spf13/cobra"
)

// testCmd represents the "mvn test" command
// If executed from the root location of pom.xml,
// it would execute the unit tests and integration tests, if any
// providing an output of the test execution on the console
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run your unit tests with maven",
	Run: func(cmd *cobra.Command, args []string) {
		runUnitTests()
	},
}

func runUnitTests() {
	msg := "Running your unit tests ... "

	// Spinner with custom message to display execution progress
	s := helper.StartSpinner(msg)

	cmd := exec.Command("mvn", "test")
	helper.PrintCommand(cmd)

	output, err := cmd.CombinedOutput()
	if err != nil {
		helper.PrintError(err)
	}

	s.Stop()
	helper.PrintOutput(output)
}

func init() {
	mavenCmd.AddCommand(testCmd)
}
