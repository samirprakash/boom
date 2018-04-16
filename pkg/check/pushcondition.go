package check

import (
	"os"
	"strings"

	"github.com/samirprakash/boom/pkg/handle"
	git "gopkg.in/src-d/go-git.v4"
)

// IfImageIsToBePushed checks if the current branch is master or develop and returns a bool
// Images should be pushed to remote repository from develop or master branches only
func IfImageIsToBePushed() bool {

	pwd, _ := os.Getwd()
	r, err := git.PlainOpen(pwd)
	handle.Error(err)
	headRef, err := r.Head()
	handle.Error(err)

	return strings.Contains(headRef.Name().String(), "develop") || strings.Contains(headRef.Name().String(), "master")
}
