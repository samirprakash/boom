package cmd

import (
	"os/exec"

	"github.com/samirprakash/go-boom/helper"
	"github.com/samirprakash/go-boom/helper/print"
	"github.com/spf13/cobra"
)

// packageCmd represents the package command
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "Package your build into a JAR -skip the tests",
	Run: func(cmd *cobra.Command, args []string) {
		generateJAR()
	},
}

func generateJAR() {
	msg := "Geneating JAR ... "

	// Spinner with custom message to display execution progress
	s := helper.StartSpinner(msg)

	cmd := exec.Command("mvn", "package", "-DskipTests")
	print.Command(cmd)

	output, err := cmd.CombinedOutput()
	if err != nil {
		print.Error(err)
	}

	s.Stop()
	print.Output(output)
}

func init() {
	mavenCmd.AddCommand(packageCmd)
}
