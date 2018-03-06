package cmd

import (
	"fmt"
	"os/exec"

	"github.com/samirprakash/boom/helpers"
	"github.com/spf13/cobra"
)

var (
	maven  bool
	gradle bool
	npm    bool
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build my code base",
	Long:  `Build the code base as I want it to and then generate an artifact and let me know where it is`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("build called")
		fmt.Printf("Maven is %v\n", maven)
		fmt.Printf("Gradle is %v\n", gradle)
		fmt.Printf("npm is %v\n", npm)

		if maven {
			cmd := exec.Command("mvn", "clean", "package")

			helpers.PrintCommand(cmd)
			output, err := cmd.CombinedOutput()
			helpers.PrintError(err)
			helpers.PrintOutput(output)
		} else if gradle {
			cmd := exec.Command("gradle", "clean", "build")

			helpers.PrintCommand(cmd)
			output, err := cmd.CombinedOutput()
			helpers.PrintError(err)
			helpers.PrintOutput(output)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().BoolVarP(&maven, "maven", "m", false, "I build your code with maven")
	buildCmd.Flags().BoolVarP(&gradle, "gradle", "g", false, "I build your code with gradle")
	buildCmd.Flags().BoolVarP(&npm, "npm", "n", false, "I build your code with node package manager")
}
