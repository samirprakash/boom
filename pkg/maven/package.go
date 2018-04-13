package maven

import (
	"github.com/samirprakash/boom/pkg/task"
)

// Package generates an artifact based on your maven based code base
func Package(flags *Flags) {
	c := "mvn package"
	if flags.SkipTests {
		c += " -DskipTests"
	}
	task.Execute(c)
}
