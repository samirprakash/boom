package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/samirprakash/go-boom/helper"
	"github.com/spf13/cobra"
)

var (
	runIntegrationTests bool
	runUnitTests        bool

	mavenCmd = &cobra.Command{
		Use:   "maven",
		Short: "Run builds with maven",
		Args:  cobra.MinimumNArgs(1),
	}

	buildCmd = &cobra.Command{
		Use:   "build",
		Short: "All in One - Validate, Compile, Clean, Test, Package, Code Coverage and Sonar",
		Run: func(cmd *cobra.Command, args []string) {
			msg := "Cleaning, running tests, pakcaging and uploading reports ... "
			// Spinner with custom message to display execution progress
			s := helper.StartSpinner(msg)
			c := exec.Command("mvn", "validate", "compile", "clean", "org.jacoco:jacoco-maven-plugin:prepare-agent", "test", "package", "-DskipTests", "sonar:sonar")
			fmt.Printf("==> Executing %s\n", strings.Join(c.Args, " "))
			output, err := c.CombinedOutput()
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("\n==> Error : %s\n", err.Error()))
			}
			s.Stop()
			if len(output) > 0 {
				fmt.Printf("==> Output : \n%s\n", string(output))
			}
		},
	}

	validateCmd = &cobra.Command{
		Use:   "validate",
		Short: "Performs a validation and checks for compilation issues",
		Run: func(cmd *cobra.Command, args []string) {
			msg := "Performing a validation and checking for compilation issues ... "
			// Spinner with custom message to display execution progress
			s := helper.StartSpinner(msg)
			c := exec.Command("mvn", "validate", "compile")
			fmt.Printf("==> Executing %s\n", strings.Join(c.Args, " "))
			output, err := c.CombinedOutput()
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("\n==> Error : %s\n", err.Error()))
			}
			s.Stop()
			if len(output) > 0 {
				fmt.Printf("==> Output : \n%s\n", string(output))
			}
		},
	}

	cleanCmd = &cobra.Command{
		Use:   "clean",
		Short: "Cleans up your workspace",
		Run: func(cmd *cobra.Command, args []string) {
			msg := "Cleaning up your build directory ... "
			// Spinner with custom message to display execution progress
			s := helper.StartSpinner(msg)
			c := exec.Command("mvn", "clean")
			fmt.Printf("==> Executing %s\n", strings.Join(c.Args, " "))
			output, err := c.CombinedOutput()
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("\n==> Error : %s\n", err.Error()))
			}
			s.Stop()
			if len(output) > 0 {
				fmt.Printf("==> Output : \n%s\n", string(output))
			}
		},
	}

	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Executes unit tests facilitated with code coverage",
		Run: func(cmd *cobra.Command, args []string) {
			msg := "Executing unit tests with code coverage ... "
			// Spinner with custom message to display execution progress
			s := helper.StartSpinner(msg)
			c := exec.Command("mvn", "org.jacoco:jacoco-maven-plugin:prepare-agent", "test")
			fmt.Printf("==> Executing %s\n", strings.Join(c.Args, " "))
			output, err := c.CombinedOutput()
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("\n==> Error : %s\n", err.Error()))
			}
			s.Stop()
			if len(output) > 0 {
				fmt.Printf("==> Output : \n%s\n", string(output))
			}
		},
	}

	packageCmd = &cobra.Command{
		Use:   "package",
		Short: "Packages your compiled code in a distributable format",
		Run: func(cmd *cobra.Command, args []string) {
			msg := "Skipping unit tests and packaging compiled code to an executable JAR ... "
			s := helper.StartSpinner(msg)
			c := exec.Command("mvn", "package", "-DskipTests")
			fmt.Printf("==> Executing %s\n", strings.Join(c.Args, " "))
			output, err := c.CombinedOutput()
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("\n==> Error : %s\n", err.Error()))
			}
			s.Stop()
			if len(output) > 0 {
				fmt.Printf("==> Output : \n%s\n", string(output))
			}
		},
	}

	verifyCmd = &cobra.Command{
		Use:   "verify",
		Short: "Runs quality checks on integration test results",
		Run: func(cmd *cobra.Command, args []string) {
			msg := "Verifying quality checks on integration test results ... "
			// Spinner with custom message to display execution progress
			s := helper.StartSpinner(msg)
			c := exec.Command("mvn", "verify")
			fmt.Printf("==> Executing %s\n", strings.Join(c.Args, " "))
			output, err := c.CombinedOutput()
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("\n==> Error : %s\n", err.Error()))
			}
			s.Stop()
			if len(output) > 0 {
				fmt.Printf("==> Output : \n%s\n", string(output))
			}
		},
	}

	deployCmd = &cobra.Command{
		Use:   "deploy",
		Short: "Copies generated package to artifactory",
		Run: func(cmd *cobra.Command, args []string) {
			msg := "Deploying generated packages to artifactory ... "
			// Spinner with custom message to display execution progress
			s := helper.StartSpinner(msg)
			c := exec.Command("mvn", "deploy")
			fmt.Printf("==> Executing %s\n", strings.Join(c.Args, " "))
			output, err := c.CombinedOutput()
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("\n==> Error : %s\n", err.Error()))
			}
			s.Stop()
			if len(output) > 0 {
				fmt.Printf("==> Output : \n%s\n", string(output))
			}
		},
	}
)

func init() {
	testCmd.Flags().BoolVarP(&runIntegrationTests, "integration-tests", "i", false, "maven test [ --integration-tests | -i ]")
	testCmd.Flags().BoolVarP(&runUnitTests, "unit-tests", "u", false, "maven test [ --unit-tests | -u ]")

	rootCmd.AddCommand(mavenCmd)

	mavenCmd.AddCommand(buildCmd)
	mavenCmd.AddCommand(validateCmd)
	mavenCmd.AddCommand(cleanCmd)
	mavenCmd.AddCommand(testCmd)
	mavenCmd.AddCommand(verifyCmd)
	mavenCmd.AddCommand(packageCmd)
	mavenCmd.AddCommand(deployCmd)
}
