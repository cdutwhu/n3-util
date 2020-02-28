package common

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// WriteFile :
func WriteFile(filename string, data []byte, perm os.FileMode) {
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0700)
	}
	FailOnErr("%v", ioutil.WriteFile(filename, data, perm))
}
