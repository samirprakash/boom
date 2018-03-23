package cmd

import (
	"github.com/spf13/cobra"
)

var (

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
	- go-boom docker build [ --image-tag | -i ] -h
	- go-boom docker compose [ --compose-file | -f ] -h
	- go-boom docker run [ --network-bridge | -n ] [ --integration-test-file | -f ] [ --environment-file | -e ] -h
	- go-boom docker tag [ --current-image | -i ] [ --new-image | -n ] -h
		`,
		Args: cobra.MinimumNArgs(1),
	}
)

func init() {
	rootCmd.AddCommand(dockerCmd)
}
