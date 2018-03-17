package cmd

import (
	"os/exec"

	"github.com/samirprakash/go-boom/helpers"
	"github.com/spf13/cobra"
)

// cleanCmd represents the "mvn clean" command
// If executed from the root location of pom.xml,
// it would remove target folders from the code base on which it is executed
// providing a clean code base for multiple and frequent builds
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Cleanup your workspace with maven",
	Run: func(cmd *cobra.Command, args []string) {
		cleanWorkingDir()
	},
}

// cleanWorkingDir would execute "mvn clean" command,
// generate error/success output and print it on console for the end user
func cleanWorkingDir() {
	msg := "Cleaning up your build directory ... "

	// Spinner with custom message to display execution progress
	s := helpers.StartSpinner(msg)

	cmd := exec.Command("mvn", "clean")
	helpers.PrintCommand(cmd)

	output, err := cmd.CombinedOutput()
	if err != nil {
		helpers.PrintError(err)
	}

	s.Stop()
	helpers.PrintOutput(output)
}

func init() {
	mavenCmd.AddCommand(cleanCmd)
}
