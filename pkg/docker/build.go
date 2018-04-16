package docker

import (
	"github.com/samirprakash/boom/pkg/check"
	"github.com/samirprakash/boom/pkg/handle"
	"github.com/samirprakash/boom/pkg/task"
)

// BuildAndPush generates and pushes docker images to remote docker registry
func BuildAndPush(flags *Flags) {
	uploadPath := flags.UploadPath
	imageTag := flags.ImageTag
	appType := flags.AppType

	if uploadPath == "" {
		handle.Info("\nMissing data - please provide the path to your docker registry. \nRun `boom docker build -h` for usage guidelines!")
		return
	} else if imageTag == "" {
		handle.Info("\nMissing data - please provide the image tag. \nRun `boom docker build -h` for usage guidelines!")
		return
	} else if appType == "" {
		handle.Info("\nMissing data - please provide the application type. \nRun `boom docker build -h` for usage guidelines!")
		return
	}

	tag := uploadPath + "/" + appType + "/" + imageTag
	buildImage := "docker build --tag " + tag + " ."
	task.Execute(buildImage)

	if check.IfImageIsToBePushed() {
		pushImage := "docker push " + tag
		task.Execute(pushImage)
	}
}
