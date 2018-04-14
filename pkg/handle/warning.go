package handle

import "fmt"

// IfWarning should be used to display a warning
// execution can proceed in case of an warnign after user has been notified
func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
