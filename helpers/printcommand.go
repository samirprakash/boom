package helpers

import (
	"fmt"
	"os/exec"
	"strings"
)

// PrintCommand represents the logic for printing the current command being executed on the console
func PrintCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing %s\n", strings.Join(cmd.Args, " "))
}
