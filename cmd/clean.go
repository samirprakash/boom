package cmd

import (
	"os/exec"

	"github.com/samirprakash/boom/helpers"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Cleanup your workspace with maven",
	Run: func(cmd *cobra.Command, args []string) {
		cleanRepo()
	},
}

func cleanRepo() {
	cmd := exec.Command("mvn", "clean")
	helpers.PrintCommand(cmd)
	output, err := cmd.CombinedOutput()
	helpers.PrintError(err)
	helpers.PrintOutput(output)
}

func init() {
	mavenCmd.AddCommand(cleanCmd)
}
