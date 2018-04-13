package maven

import (
	"github.com/samirprakash/boom/pkg/task"
)

// Test executes integration and unit tests for your maven based code base
func Test(flags *Flags) {
	c := "mvn org.jacoco:jacoco-maven-plugin:prepare-agent test"
	switch {
	case flags.RunIntegrationTests:
		c += " -Dcategories=integration-tests"
	case flags.RunUnitTests:
		c += " -Dcategories=unit-tests"
	}
	task.Execute(c)
}
