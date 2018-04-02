package maven

import (
	"strconv"

	"github.com/samirprakash/boom/pkg/task"
	"github.com/spf13/cobra"
)

// Test executes integration and unit tests for your maven based code base
func Test(cmd *cobra.Command, args []string) {
	c := "mvn org.jacoco:jacoco-maven-plugin:prepare-agent test"
	runIntegrationTests, _ := strconv.ParseBool(args[0])
	runUnitTests, _ := strconv.ParseBool(args[1])
	switch {
	case runIntegrationTests:
		c += " -Dcategories=integration-tests"
	case runUnitTests:
		c += " -Dcategories=unit-tests"
	}
	task.Execute(c)
}
