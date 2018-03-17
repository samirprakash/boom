package helpers

import "fmt"

// PrintOutput represents the logic for printing the final output of the command that has been executed
func PrintOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output : \n%s\n", string(outs))
	}
}
