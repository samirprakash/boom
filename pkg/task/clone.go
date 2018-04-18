package task

import (
	"os"

	log "github.com/sirupsen/logrus"
	git "gopkg.in/src-d/go-git.v4"
)

// Clone would clone a git repository to the specified path from the specified repoName
func Clone(path, repoName string) {
	uname := os.Getenv("GIT_USERNAME")
	pwd := os.Getenv("GIT_PASSWORD")

	url := "https://" + uname + ":" + pwd + "@github.com/" + repoName + ".git"
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	if err != nil {
		log.Fatal(err)
	}
}
