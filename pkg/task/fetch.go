package task

import (
	log "github.com/sirupsen/logrus"
	git "gopkg.in/src-d/go-git.v4"
)

// Fetch would pull latest changes from the origin configured to the git repository specified in the path variable
func Fetch(path string) {
	r, err := git.PlainOpen(path)
	if err != nil {
		log.Fatal(err)
	}
	w, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		log.Warn(err.Error())
	}
}
