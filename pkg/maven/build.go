package maven

import (
	"github.com/samirprakash/boom/pkg/task"
)

// Build does the full fledged build for your maven based code base
func Build() {
	c := "mvn validate compile clean org.jacoco:jacoco-maven-plugin:prepare-agent test package -DskipTests sonar:sonar"
	task.Execute(c)
}
