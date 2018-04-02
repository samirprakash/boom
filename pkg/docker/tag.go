package docker

import (
	"fmt"
	"os"

	"github.com/samirprakash/boom/pkg/task"
	"github.com/spf13/cobra"
)

// TagAndPush tags the exisitng images and pushes them to a remote docker registry
func TagAndPush(cmd *cobra.Command, args []string) {
	currentImage := args[0]
	newImage := args[1]
	if currentImage == "" {
		fmt.Fprintln(os.Stderr, "\nMissing data - please provide the current image tag. \nRun `boom docker tag -h` for usage guidelines!")
		return
	} else if newImage == "" {
		fmt.Fprintln(os.Stderr, "\nMissing data - please provide the new image tag. \nRun `boom docker tag -h` for usage guidelines!")
		return
	}

	c := "docker tag " + currentImage + " " + newImage
	task.Execute(c)
	c = "docker push " + newImage
	task.Execute(c)
}
