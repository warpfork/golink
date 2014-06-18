package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// For every case where these functions might fail during golink, there can be no recovery.
// Wrap stdlib functions to panic when failed.

func CreateFolder(folder string) {
	err := os.MkdirAll(folder, 0755)
	if err != nil { ExitGently(err) }
}

func Symlink(destination, source string) {
	err := os.Symlink(destination, source)
	if err != nil { ExitGently(err) }
}

func WriteFile(filename, data string, mode os.FileMode) {
	str := []byte(data)
	err := ioutil.WriteFile(filename, str, mode)
	if err != nil { ExitGently(err) }
}

func Abs(path string) string {
	result, err := filepath.Abs(path)
	if err != nil { ExitGently(err) }
	return result
}
