package task

import (
	"github.com/samirprakash/boom/pkg/handle"
	git "gopkg.in/src-d/go-git.v4"
)

// Fetch would pull latest changes from the origin configured to the git repository specified in the path variable
func Fetch(path string) {
	r, err := git.PlainOpen(path)
	handle.Error(err)
	w, err := r.Worktree()
	handle.Error(err)
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		handle.Warning(err.Error())
	}
}
