package check

import "os"

// IfDirExists checks for a directory to be present
func IfDirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		return false, nil
	}
	if os.IsNotExist(err) {
		return true, nil
	}
	return true, err
}
