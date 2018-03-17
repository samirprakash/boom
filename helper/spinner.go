package helper

import (
	"time"

	"github.com/briandowns/spinner"
)

// StartSpinner represents the logic for creating a new spinner with a custom message
// Finally, it returns back a pointer to spinner to be used from the calling func
// The return value would be used to stop the spinner once the actualt command execution has completed
func StartSpinner(msg string) (s *spinner.Spinner) {
	s = spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Prefix = msg
	s.Color("green")
	s.Start()

	return s
}
