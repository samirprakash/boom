package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mavenCmd represents the maven command
var mavenCmd = &cobra.Command{
	Use:   "maven",
	Short: "Build with maven",
	Long: `This is the base command for running builds with maven dependency management:

- Clean, build, test, package and generate sonar reports
- Clean your workspace
- Build your workspace
- Test your workspace
- Generate sonar repots`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("maven called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(mavenCmd)
}
