package maven

import (
	"github.com/samirprakash/boom/utils"
	"github.com/spf13/cobra"
)

// Build does the full fledged build for your maven based code base
func Build(cmd *cobra.Command, args []string) {
	c := "mvn validate compile clean org.jacoco:jacoco-maven-plugin:prepare-agent test package -DskipTests sonar:sonar"
	utils.Execute(c)
}
