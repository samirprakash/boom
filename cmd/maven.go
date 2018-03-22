package cmd

import (
	"os"
	"os/exec"
	"strings"

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
			c := "mvn validate compile clean org.jacoco:jacoco-maven-plugin:prepare-agent test package -DskipTests sonar:sonar"
			execute(c)
		},
	}

	validateCmd = &cobra.Command{
		Use:   "validate",
		Short: "Performs a validation and checks for compilation issues",
		Run: func(cmd *cobra.Command, args []string) {
			c := "mvn validate compile"
			execute(c)
		},
	}

	cleanCmd = &cobra.Command{
		Use:   "clean",
		Short: "Cleans up your workspace",
		Run: func(cmd *cobra.Command, args []string) {
			c := "mvn clean"
			execute(c)
		},
	}

	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Executes unit tests facilitated with code coverage",
		Run: func(cmd *cobra.Command, args []string) {
			c := "mvn org.jacoco:jacoco-maven-plugin:prepare-agent test"
			execute(c)
		},
	}

	packageCmd = &cobra.Command{
		Use:   "package",
		Short: "Packages your compiled code in a distributable format",
		Run: func(cmd *cobra.Command, args []string) {
			c := "mvn package -DskipTests"
			execute(c)
		},
	}

	verifyCmd = &cobra.Command{
		Use:   "verify",
		Short: "Runs quality checks on integration test results",
		Run: func(cmd *cobra.Command, args []string) {
			c := "mvn verify"
			execute(c)
		},
	}

	deployCmd = &cobra.Command{
		Use:   "deploy",
		Short: "Copies generated package to artifactory",
		Run: func(cmd *cobra.Command, args []string) {
			c := "mvn deploy"
			execute(c)
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

func execute(c string) {
	tokens := strings.Fields(c)
	executable := tokens[0]
	tokens = tokens[1:len(tokens)]

	cmd := exec.Command(executable, tokens...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		os.Exit(2)
	}
}
