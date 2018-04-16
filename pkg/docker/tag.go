package docker

import (
	"github.com/samirprakash/boom/pkg/check"
	"github.com/samirprakash/boom/pkg/handle"
	"github.com/samirprakash/boom/pkg/task"
)

// TagAndPush tags the exisitng images and pushes them to a remote docker registry
func TagAndPush(flags *Flags) {
	currentImage := flags.CurrentImage
	newImage := flags.NewImage
	if currentImage == "" {
		handle.Info("\nMissing data - please provide the current image tag. \nRun `boom docker tag -h` for usage guidelines!")
		return
	} else if newImage == "" {
		handle.Info("\nMissing data - please provide the new image tag. \nRun `boom docker tag -h` for usage guidelines!")
		return
	}

	c := "docker tag " + currentImage + " " + newImage
	task.Execute(c)

	if check.IfImageIsToBePushed() {
		c = "docker push " + newImage
		task.Execute(c)
	}
}
