package print

import (
	"fmt"
	"os"
)

// Error represents the logic for printing the errors generated from the CLI execution, if any
func Error(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("\n==> Error : %s\n", err.Error()))
	}
}
