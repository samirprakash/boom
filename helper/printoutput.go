package helper

import "fmt"

// PrintOutput represents the logic for printing the final output of the command that has been executed
func PrintOutput(output []byte) {
	if len(output) > 0 {
		fmt.Printf("==> Output : \n%s\n", string(output))
	}
}
