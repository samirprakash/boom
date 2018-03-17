package cmd

import (
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/samirprakash/go-boom/helpers"
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
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Prefix = "Cleaning up your build directory ... "
	s.Color("green")
	s.Start()

	cmd := exec.Command("mvn", "clean")
	output, err := cmd.CombinedOutput()
	s.Stop()

	helpers.PrintCommand(cmd)
	helpers.PrintError(err)
	helpers.PrintOutput(output)
}

func init() {
	mavenCmd.AddCommand(cleanCmd)
}
