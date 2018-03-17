package helpers

import (
	"fmt"
	"os"
)

//PrintError prints the errors generated from the CLI execution, if any
func PrintError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("\n==> Error : %s\n", err.Error()))
	}
}
