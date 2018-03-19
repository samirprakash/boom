package cmd

import (
	"os/exec"

	"github.com/samirprakash/go-boom/helper"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
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
