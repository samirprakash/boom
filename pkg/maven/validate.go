package maven

import (
	"github.com/samirprakash/boom/pkg/task"
)

// Validate compiles and validates your maveb based code base
func Validate() {
	c := "mvn validate compile"
	task.Execute(c)
}
