package docker

import (
	"fmt"
	"os"

	"github.com/samirprakash/boom/pkg/task"
)

// BuildAndPush generates and pushes docker images to remote docker registry
func BuildAndPush(flags *Flags) {
	uploadPath := flags.UploadPath
	imageTag := flags.ImageTag
	appType := flags.AppType

	if uploadPath == "" {
		fmt.Fprintln(os.Stderr, "\nMissing data - please provide the path to your docker registry. \nRun `boom docker build -h` for usage guidelines!")
		return
	} else if imageTag == "" {
		fmt.Fprintln(os.Stderr, "\nMissing data - please provide the image tag. \nRun `boom docker build -h` for usage guidelines!")
		return
	} else if appType == "" {
		fmt.Fprintln(os.Stderr, "\nMissing data - please provide the application type. \nRun `boom docker build -h` for usage guidelines!")
		return
	}
	tag := uploadPath + "/" + appType + "/" + imageTag
	buildImage := "docker build --tag " + tag + " ."
	pushImage := "docker push " + tag
	task.Execute(buildImage)
	task.Execute(pushImage)
}
