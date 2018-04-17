package maven

import (
	"github.com/samirprakash/boom/pkg/task"
)

// Verify defines the option to validate and compile maven based code base
func Verify() {
	c := "mvn verify"
	task.Execute(c)
}
