package cmd

import "os"

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		return false, nil
	}
	if os.IsNotExist(err) {
		return true, nil
	}
	return true, err
}
