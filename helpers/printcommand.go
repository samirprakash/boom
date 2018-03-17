package helpers

import (
	"fmt"
	"os/exec"
	"strings"
)

// PrintCommand Export option to print commands
func PrintCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing %s\n", strings.Join(cmd.Args, " "))
}
