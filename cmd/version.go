package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display boom version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(rootCmd.Use + " v" + VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
