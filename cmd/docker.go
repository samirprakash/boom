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
	- go-boom docker run [ --network-bridge | -n ] [ --integration-test-file | -f ] [ --environment-file | -e ] -h
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
			cloneConfig := "git clone git@github.com:toyota-connected/pg-config-source.git " + os.Getenv("TC_CONFIG_PATH")
			setupEnvironment := "docker-compose -f " + composeFile + " up --build --detach --remove-orphans"

			fmt.Println(cloneConfig)
			// execute(cloneConfig)
			fmt.Println(setupEnvironment)
			execute(setupEnvironment)
			healthcheck(healthcheckPorts)
			// buildImage := "docker build --tag " + tag + " ."
			// pushImage := "docker push " + tag
			// execute(buildImage)
			// execute(pushImage)
		},
	}
)

func init() {
	imageCmd.Flags().StringVarP(&uploadPath, "upload-to", "u", "", "specify the url to your docker registry")
	imageCmd.Flags().StringVarP(&imageTag, "image-tag", "i", "", "specify the tag for your image")
	imageCmd.Flags().StringVarP(&appType, "app-type", "t", "", "specifcy the application type - services/client")
	composeCmd.Flags().StringVarP(&composeFile, "compose-file", "f", "", "specify the compose file to used for setting up the environment")
	composeCmd.Flags().StringVarP(&healthcheckPorts, "healthcheck-ports", "p", "", "specify the healthcheck ports exposed in the compose file - use comma seperated format")

	rootCmd.AddCommand(dockerCmd)
	dockerCmd.AddCommand(imageCmd)
	dockerCmd.AddCommand(composeCmd)
}
