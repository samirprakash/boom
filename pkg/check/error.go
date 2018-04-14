package check

import (
	"fmt"
	"os"
)

// IfError should be used to naively panic if an error is not nil.
func IfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
