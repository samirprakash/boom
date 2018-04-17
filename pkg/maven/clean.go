package maven

import (
	"github.com/samirprakash/boom/pkg/task"
)

// Clean removes existing target folder from your maven based code base
func Clean() {
	c := "mvn clean"
	task.Execute(c)
}
