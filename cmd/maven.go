package cmd

import (
	"fmt"
	"os"

	"github.com/samirprakash/boom/utils"
	"github.com/spf13/cobra"
)

var (
	runIntegrationTests bool
	runUnitTests        bool
	skipTests           bool
	repoID              string

	// mavenCmd is the parent command to execute maven build steps
	// execute boom maven -h to check the available options
	mavenCmd = &cobra.Command{
		Use:   "maven",
		Short: "Run builds with maven",
		Long: `
[ boom maven ] provides an option to execute basic maven build steps.
It requires at least one sub command from the list of options to be specified.

Prerequisites:
	- Install this binary and add it to your path
	- Install maven locally
	- Keep a settings.xml file in your .m2 folder

Example usage options:
	- boom maven [ build | validate | clean | test | package | verify | deploy ] -h
	- boom maven build
	- boom maven validate
	- boom maven clean
	- boom maven test [ -integration-tests | -i ] [ --unit-tests | -u ]
	- boom maven package [ --skip-test | -s ]
	- boom maven verify
	- boom maven deploy [ --repository-id ]
		`,
		Args: cobra.MinimumNArgs(1),
	}

	// buildCmd is the subcommand to execute all the maven build steps in one go for your maven based code base
	buildCmd = &cobra.Command{
		Use:     "build",
		Short:   "All in One - Validate, Compile, Clean, Unit test, Package, Code Coverage and Sonar",
		Example: "boom maven build -h",
		Run:     builder,
	}

	// validateCmd is the subcommand to validate and compile your maven based code base
	validateCmd = &cobra.Command{
		Use:     "validate",
		Short:   "Performs a validation and checks for compilation issues",
		Example: "boom maven validate -h",
		Run:     validator,
	}

	// cleanCmd is the subcommand to clean your maven based code base
	cleanCmd = &cobra.Command{
		Use:     "clean",
		Short:   "Cleans up your workspace",
		Example: "boom maven clean -h",
		Run:     cleaner,
	}

	// testCmd is the subcommand to execute unit tests in your maven based code base
	// this command executes only the integration tests if the [ --integration-tests | -i ] has been set
	// this command executes only the unit tests if the [ --unit-tests | -u ] has been set
	testCmd = &cobra.Command{
		Use:     "test",
		Short:   "Executes unit tests facilitated with code coverage",
		Example: "boom maven test [ --integration-tests | -i ] [ --unit-tests | -u ] -h",
		Run:     tester,
	}

	// packageCmd is the subcommand to package your maven based code base
	// this command skips the tests if the [ --skip-tests | -s ] flag has been set
	packageCmd = &cobra.Command{
		Use:     "package",
		Short:   "Packages your compiled code in a distributable format",
		Example: "boom maven package [ --skip-tests | -s ] -h",
		Run:     packager,
	}

	// verifyCmd is the subcommand to verify the integration test results after the maven based code base has been packaged
	verifyCmd = &cobra.Command{
		Use:     "verify",
		Short:   "Runs quality checks on integration test results",
		Example: "boom maven verify -h",
		Run:     verifier,
	}

	// deployCmd is the subcommand to deploy your packaged maven based code base to remote repository
	// remote repository location can be specified using [ --repository-id ] flag with the remote repository id specified in your local `.m2/settings.xml`
	deployCmd = &cobra.Command{
		Use:     "deploy",
		Short:   "Copies generated package to artifactory or nexus",
		Example: "boom maven deploy --repository-id {your-repo-id-defined-in-your-maven-settings} -h",
		Run:     deployer,
	}
)

func init() {
	// Add flags to the sub commands to logical selection of options
	testCmd.Flags().BoolVarP(&runIntegrationTests, "integration-tests", "i", false, "Use this flag to execute integration tests")
	testCmd.Flags().BoolVarP(&runUnitTests, "unit-tests", "u", false, "Use this flag to execute unit tests")

	packageCmd.Flags().BoolVarP(&skipTests, "skip-tests", "s", false, "Use this flag to skip test while packaging")

	deployCmd.Flags().StringVar(&repoID, "repository-id", "", "Provide this value to connect to the remote repository. Value must be from local .m2/settings.xml")

	rootCmd.AddCommand(mavenCmd)
	mavenCmd.AddCommand(buildCmd, validateCmd, cleanCmd, testCmd, verifyCmd, packageCmd, deployCmd)
}

func builder(cmd *cobra.Command, args []string) {
	c := "mvn validate compile clean org.jacoco:jacoco-maven-plugin:prepare-agent test package -DskipTests sonar:sonar"
	utils.Execute(c)
}

func validator(cmd *cobra.Command, args []string) {
	c := "mvn validate compile"
	utils.Execute(c)
}

func cleaner(cmd *cobra.Command, args []string) {
	c := "mvn clean"
	utils.Execute(c)
}

func tester(cmd *cobra.Command, args []string) {
	c := "mvn org.jacoco:jacoco-maven-plugin:prepare-agent test"
	switch {
	case runIntegrationTests:
		c += " -Dcategories=integration-tests"
	case runUnitTests:
		c += " -Dcategories=unit-tests"
	}
	utils.Execute(c)
}

func packager(cmd *cobra.Command, args []string) {
	c := "mvn package"
	if skipTests {
		c += " -DskipTests"
	}
	utils.Execute(c)
}

func verifier(cmd *cobra.Command, args []string) {
	c := "mvn verify"
	utils.Execute(c)
}

func deployer(cmd *cobra.Command, args []string) {
	if repoID == "" {
		fmt.Fprintln(os.Stderr, "Missing data - please provide the repository id. \n\nRun `boom maven deploy -h` for usage guidelines!")
		return
	}
	c := "mvn deploy -DrepositoryId=" + repoID
	utils.Execute(c)
}
