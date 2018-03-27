package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	imageTag         string
	appType          string
	uploadPath       string
	composeFile      string
	healthcheckPorts string
	networkBridge    string
	testCollection   string
	environmentSpec  string
	currentImage     string
	newImage         string

	// dockerCmd is the parent command to execute docker and docker-compose actions
	// execute go-boom docker -h to check the available options
	dockerCmd = &cobra.Command{
		Use:   "docker",
		Short: "Execute docker commands",
		Long: `
[ go-boom docker ] provides an option to execute basic docker commands.
It requires at least one sub command from the list of options to be specified.

Prerequisites:
	- Install this binary and add it to your path
	- Install docker

Example usage options:
	- go-boom docker [ build | compose | run | tag ] -h
	- go-boom docker build [ --image-tag | -i ] [ --app-type | -t ] -h
	- go-boom docker compose [ --compose-file | -f ] -h
	- go-boom docker run [ --network-bridge | -n ] [ --test-collection | -c ] [ --environment-spec | -e ] -h
	- go-boom docker tag [ --current-image | -i ] [ --new-image | -n ] -h
		`,
		Args: cobra.MinimumNArgs(1),
	}

	// imageCmd is the subcommand to generate docker images
	// execute go-boom docker build -h to check the available options
	imageCmd = &cobra.Command{
		Use:     "build",
		Short:   "Build docker images and push to a remote repository",
		Example: "go-boom docker build [ --image-tag | -i ] [ --app-type | -t ] -h",
		Run: func(cmd *cobra.Command, args []string) {
			if uploadPath == "" {
				fmt.Fprintln(os.Stderr, "\nMissing data - please provide the path to your docker registry. \nRun `go-boom docker build -h` for usage guidelines!")
				return
			} else if imageTag == "" {
				fmt.Fprintln(os.Stderr, "\nMissing data - please provide the image tag. \nRun `go-boom docker build -h` for usage guidelines!")
				return
			} else if appType == "" {
				fmt.Fprintln(os.Stderr, "\nMissing data - please provide the application type. \nRun `go-boom docker build -h` for usage guidelines!")
				return
			}
			tag := uploadPath + "/" + appType + "/" + imageTag
			buildImage := "docker build --tag " + tag + " ."
			pushImage := "docker push " + tag
			execute(buildImage)
			execute(pushImage)
		},
	}

	// compose is the subcommand to start a docker compose environment to integration testing
	// execute go-boom docker compose -h to check the available options
	composeCmd = &cobra.Command{
		Use:     "compose",
		Short:   "Create docker compose environment based on the docker-compose.yaml in the code base",
		Example: "go-boom docker compose [ --compose-file | -f ] [ --healthcheck-ports | -p ] -h",
		Run: func(cmd *cobra.Command, args []string) {
			if composeFile == "" {
				fmt.Fprintln(os.Stderr, "\nMissing data - please provide the docker compose file. \nRun `go-boom docker compose -h` for usage guidelines!")
				return
			} else if healthcheckPorts == "" {
				fmt.Fprintln(os.Stderr, "\nMissing data - please provide the healthcheck ports exposed in the docker compose file. \nRun `go-boom docker compose -h` for usage guidelines!")
				return
			}

			// clone config source repo if not already present in the build environment
			path := os.Getenv("TC_CONFIG_PATH")
			repo, _ := exists(path)
			if !repo {
				fmt.Println("cloning into : ", path)
				cloneConfig := "git clone git@github.com:toyota-connected/pg-config-source.git " + path
				execute(cloneConfig)
			}
			fmt.Println("repository that is being cloned already exists on the build environment")

			// setup docker compose environment based on the specified docker compose file
			setupEnvironment := "docker-compose -f " + composeFile + " up --build --detach --remove-orphans"
			execute(setupEnvironment)

			// check if the docker containers are healthy or not based on the ports that have been exposed from docker-compose.yaml
			healthcheck(healthcheckPorts)
		},
	}

	// run is the subcommand to execute tests collection on an existing docker compose environment
	// this should be executed after `go-boom docker compose` has been executed successfully
	// execute go-boom docker run -h to check the available options
	runCmd = &cobra.Command{
		Use:     "run",
		Short:   "run collection of tests using newman command line runner",
		Example: "go-boom docker run [ --network-bridge | -n ] [ --test-collection | -c ] [ --environment-spec | -e ] -h",
		Run: func(cmd *cobra.Command, args []string) {
			if networkBridge == "" {
				fmt.Fprintln(os.Stderr, "\nMissing data - please provide network bridge name. \nRun `docker network ls` to get a list of existing network bridges!")
				fmt.Fprintln(os.Stderr, "\nIf network bridge does not exist then execute `go-boom docker compose` before running this command!")
				fmt.Fprintln(os.Stderr, "\nRun `go-boom docker run -h` for usage guidelines!")
				return
			} else if testCollection == "" {
				fmt.Fprintln(os.Stderr, "\nMissing data - please provide the JSON file name for your test collection. \nRun `go-boom docker run -h` for usage guidelines!")
				return
			} else if environmentSpec == "" {
				fmt.Fprintln(os.Stderr, "\nMissing data - please provide the JSON file defining your execution environment. \nRun `go-boom docker run -h` for usage guidelines!")
				return
			}

			pwd, _ := os.Getwd()
			v := pwd + "/integration-tests:/etc/postman postman/newman_alpine33:3.8.3"
			c := "/etc/postman/" + testCollection
			e := "/etc/postman/" + environmentSpec
			runTests := "docker run --network " + networkBridge + " -v " + v + " -c=" + c + " -e=" + e
			execute(runTests)
		},
	}

	// tag is the subcommand to tag and push images created by `go-doom docker compose` command
	// execute go-boom docker tag -h to check the available options
	tagCmd = &cobra.Command{
		Use:     "tag",
		Short:   "tag and push images to docker registry",
		Example: "go-boom docker tag [ --current-image | -i ] [ --new-image | -n ] -h",
		Run: func(cmd *cobra.Command, args []string) {
			if currentImage == "" {
				fmt.Fprintln(os.Stderr, "\nMissing data - please provide the current image tag. \nRun `go-boom docker tag -h` for usage guidelines!")
				return
			} else if newImage == "" {
				fmt.Fprintln(os.Stderr, "\nMissing data - please provide the new image tag. \nRun `go-boom docker tag -h` for usage guidelines!")
				return
			}

			c := "docker tag " + currentImage + " " + newImage
			execute(c)
			c = "docker push " + newImage
			execute(c)
		},
	}
)

func init() {
	imageCmd.Flags().StringVarP(&uploadPath, "upload-to", "u", "", "specify the url to your docker registry")
	imageCmd.Flags().StringVarP(&imageTag, "image-tag", "i", "", "specify the tag for your image")
	imageCmd.Flags().StringVarP(&appType, "app-type", "t", "", "specifcy the application type - services/client")

	composeCmd.Flags().StringVarP(&composeFile, "compose-file", "f", "", "specify the compose file to used for setting up the environment")
	composeCmd.Flags().StringVarP(&healthcheckPorts, "healthcheck-ports", "p", "", "specify the healthcheck ports exposed in the compose file - use comma seperated format")

	runCmd.Flags().StringVarP(&networkBridge, "network-bridge", "n", "", "specify the network briidge applicable for running these tests")
	runCmd.Flags().StringVarP(&testCollection, "test-collection", "c", "", "specify the test collection file name in your integration-tests folder")
	runCmd.Flags().StringVarP(&environmentSpec, "environment-file", "e", "", "specify the newman environment file name in your integration-tests folder")

	tagCmd.Flags().StringVarP(&currentImage, "current-image", "i", "", "specify the tag of existing docker image")
	tagCmd.Flags().StringVarP(&newImage, "new-image", "n", "", "specify the tag name to tag the existing image with")

	rootCmd.AddCommand(dockerCmd)
	dockerCmd.AddCommand(imageCmd)
	dockerCmd.AddCommand(composeCmd)
	dockerCmd.AddCommand(runCmd)
	dockerCmd.AddCommand(tagCmd)
}
