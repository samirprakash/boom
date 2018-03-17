package cmd

import (
	"github.com/spf13/cobra"
)

// mavenCmd represents the root command to execute CLI tasks with maven dependencies
var mavenCmd = &cobra.Command{
	Use:   "maven",
	Short: "Run builds with maven",
	Long: `This is the base command for running builds with maven dependency management:

- Clean, build, test, package and generate sonar reports
- Clean your workspace
- Build your workspace
- Test your workspace
- Generate sonar repots`,
	Args: cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(mavenCmd)
}
