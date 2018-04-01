package utils

import "os"

// Exists checks for a directory to be present
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		return false, nil
	}
	if os.IsNotExist(err) {
		return true, nil
	}
	return true, err
}
