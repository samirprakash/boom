package docker

import (
	"github.com/samirprakash/boom/pkg/check"
	"github.com/samirprakash/boom/pkg/task"
	log "github.com/sirupsen/logrus"
)

// BuildAndPush generates and pushes docker images to remote docker registry
func BuildAndPush(flags *Flags) {
	uploadPath := flags.UploadPath
	imageTag := flags.ImageTag
	appType := flags.AppType

	if uploadPath == "" {
		log.Fatal("\nMissing data - please provide the path to your docker registry. \nRun `boom docker build -h` for usage guidelines!")
	} else if imageTag == "" {
		log.Fatal("\nMissing data - please provide the image tag. \nRun `boom docker build -h` for usage guidelines!")
	} else if appType == "" {
		log.Fatal("\nMissing data - please provide the application type. \nRun `boom docker build -h` for usage guidelines!")
	}

	tag := uploadPath + "/" + appType + "/" + imageTag
	buildImage := "docker build --tag " + tag + " ."
	task.Execute(buildImage)

	if check.IsBranchMorD() {
		pushImage := "docker push " + tag
		task.Execute(pushImage)
	}
}
