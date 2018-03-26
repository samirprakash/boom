package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	imageTag   string
	appType    string
	uploadPath string

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
	- go-boom docker build [ --image-tag | -i ] [[ --app-type | -t ] -h
	- go-boom docker compose [ --compose-file | -f ] -h
	- go-boom docker run [ --network-bridge | -n ] [ --integration-test-file | -f ] [ --environment-file | -e ] -h
	- go-boom docker tag [ --current-image | -i ] [ --new-image | -n ] -h
		`,
		Args: cobra.MinimumNArgs(1),
	}

	// buildCmd is the subcommand to generate docker images
	// execute go-boom docker build -h to check the available options
	builderCmd = &cobra.Command{
		Use:     "build",
		Short:   "Build docker images and push to a remote repository",
		Example: "go-boom docker build [ --image-tag | -i ] [[ --app-type | -t ] -h",
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
)

func init() {
	builderCmd.Flags().StringVarP(&uploadPath, "upload-to", "u", "", "specify the url to your docker registry")
	builderCmd.Flags().StringVarP(&imageTag, "image-tag", "i", "", "specify the tag for your image")
	builderCmd.Flags().StringVarP(&appType, "app-type", "t", "", "specifcy the application type - services/client")

	rootCmd.AddCommand(dockerCmd)
	dockerCmd.AddCommand(builderCmd)
}
