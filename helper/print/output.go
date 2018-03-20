package print

import "fmt"

// Output represents the logic for printing the final output of the command that has been executed
func Output(output []byte) {
	if len(output) > 0 {
		fmt.Printf("==> Output : \n%s\n", string(output))
	}
}
