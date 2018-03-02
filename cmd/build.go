package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

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

			printCommand(cmd)
			output, err := cmd.CombinedOutput()
			printError(err)
			printOutput(output)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().BoolVarP(&maven, "maven", "m", false, "I build your code with maven")
	buildCmd.Flags().BoolVarP(&gradle, "gradle", "g", false, "I build your code with gradle")
	buildCmd.Flags().BoolVarP(&npm, "npm", "n", false, "I build your code with node package manager")
}

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error : %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output : %s\n", string(outs))
	}
}
