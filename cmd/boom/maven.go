package main

import (
	"github.com/samirprakash/boom/pkg/maven"
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
	- boom maven build -h
	- boom maven validate -h
	- boom maven clean -h
	- boom maven test [ --integration-tests | -i ] [ --unit-tests | -u ] -h
	- boom maven package [ --skip-test | -s ] -h
	- boom maven verify -h
	- boom maven deploy [ --repository-id ] -h
		`,
		Args: cobra.MinimumNArgs(1),
	}

	// buildCmd is the subcommand to execute all the maven build steps in one go for your maven based code base
	buildCmd = &cobra.Command{
		Use:     "build",
		Short:   "All in One - Validate, Compile, Clean, Unit test, Package, Code Coverage and Sonar",
		Example: "boom maven build -h",
		Run:     maven.Build,
	}

	// validateCmd is the subcommand to validate and compile your maven based code base
	validateCmd = &cobra.Command{
		Use:     "validate",
		Short:   "Performs a validation and checks for compilation issues",
		Example: "boom maven validate -h",
		Run:     maven.Validate,
	}

	// cleanCmd is the subcommand to clean your maven based code base
	cleanCmd = &cobra.Command{
		Use:     "clean",
		Short:   "Cleans up your workspace",
		Example: "boom maven clean -h",
		Run:     maven.Clean,
	}

	// testCmd is the subcommand to execute unit tests in your maven based code base
	// this command executes only the integration tests if the [ --integration-tests | -i ] has been set
	// this command executes only the unit tests if the [ --unit-tests | -u ] has been set
	testCmd = &cobra.Command{
		Use:     "test",
		Short:   "Executes unit tests facilitated with code coverage",
		Example: "boom maven test [ --integration-tests | -i ] [ --unit-tests | -u ] -h",
		Run: func(cmd *cobra.Command, args []string) {
			flags := maven.Flags{
				RunIntegrationTests: runIntegrationTests,
				RunUnitTests:        runUnitTests,
			}
			maven.Test(&flags)
		},
	}

	// packageCmd is the subcommand to package your maven based code base
	// this command skips the tests if the [ --skip-tests | -s ] flag has been set
	packageCmd = &cobra.Command{
		Use:     "package",
		Short:   "Packages your compiled code in a distributable format",
		Example: "boom maven package [ --skip-tests | -s ] -h",
		Run: func(cmd *cobra.Command, args []string) {
			flags := maven.Flags{
				SkipTests: skipTests,
			}
			maven.Package(&flags)
		},
	}

	// verifyCmd is the subcommand to verify the integration test results after the maven based code base has been packaged
	verifyCmd = &cobra.Command{
		Use:     "verify",
		Short:   "Runs quality checks on integration test results",
		Example: "boom maven verify -h",
		Run:     maven.Verify,
	}

	// deployCmd is the subcommand to deploy your packaged maven based code base to remote repository
	// remote repository location can be specified using [ --repository-id ] flag with the remote repository id specified in your local `.m2/settings.xml`
	deployCmd = &cobra.Command{
		Use:     "deploy",
		Short:   "Copies generated package to artifactory or nexus",
		Example: "boom maven deploy --repository-id {your-repo-id-defined-in-your-maven-settings} -h",
		Run: func(cmd *cobra.Command, args []string) {
			flags := maven.Flags{
				RepoID: repoID,
			}
			maven.Deploy(&flags)
		},
	}
)

func init() {
	testCmd.Flags().BoolVarP(&runIntegrationTests, "integration-tests", "i", false, "Use this flag to execute integration tests")
	testCmd.Flags().BoolVarP(&runUnitTests, "unit-tests", "u", false, "Use this flag to execute unit tests")
	packageCmd.Flags().BoolVarP(&skipTests, "skip-tests", "s", false, "Use this flag to skip test while packaging")
	deployCmd.Flags().StringVar(&repoID, "repository-id", "", "Provide this value to connect to the remote repository. Value must be from local .m2/settings.xml")

	rootCmd.AddCommand(mavenCmd)
	mavenCmd.AddCommand(buildCmd, validateCmd, cleanCmd, testCmd, verifyCmd, packageCmd, deployCmd)
}
