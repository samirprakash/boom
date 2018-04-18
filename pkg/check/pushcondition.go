package check

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	git "gopkg.in/src-d/go-git.v4"
)

// IsBranchMorD checks if the current branch is master or develop and returns a bool
// Images should be pushed to remote repository from develop or master branches only
func IsBranchMorD() bool {
	pwd, _ := os.Getwd()

	r, err := git.PlainOpen(pwd)
	if err != nil {
		log.Fatal(err)

	}
	headRef, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}

	return strings.Contains(headRef.Name().String(), "develop") || strings.Contains(headRef.Name().String(), "master")
}
