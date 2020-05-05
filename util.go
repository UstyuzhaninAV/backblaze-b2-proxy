package main

// check if file exists
import (
	"os"
)

// file check handler
func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
