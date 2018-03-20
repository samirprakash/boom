package print

import (
	"fmt"
	"os/exec"
	"strings"
)

// Command represents the logic for printing the current command being executed on the console
func Command(cmd *exec.Cmd) {
	fmt.Printf("==> Executing %s\n", strings.Join(cmd.Args, " "))
}
