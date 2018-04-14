package handle

import "fmt"

// IfInfo should be used to describe the example commands that are about to run
// or to display any information on the console.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
