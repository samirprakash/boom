package docker

import (
	"github.com/samirprakash/boom/pkg/check"
	"github.com/samirprakash/boom/pkg/task"
	log "github.com/sirupsen/logrus"
)

// TagAndPush tags the exisitng images and pushes them to a remote docker registry
func TagAndPush(flags *Flags) {
	currentImage := flags.CurrentImage
	newImage := flags.NewImage
	if currentImage == "" {
		log.Fatal("\nMissing data - please provide the current image tag. \nRun `boom docker tag -h` for usage guidelines!")
	} else if newImage == "" {
		log.Fatal("\nMissing data - please provide the new image tag. \nRun `boom docker tag -h` for usage guidelines!")
	}

	c := "docker tag " + currentImage + " " + newImage
	task.Execute(c)

	if check.IsBranchMorD() {
		c = "docker push " + newImage
		task.Execute(c)
	}
}
