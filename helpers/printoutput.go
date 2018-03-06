package helpers

import "fmt"

// PrintOutput prints the output of CLI execution
func PrintOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output : %s\n", string(outs))
	}
}
