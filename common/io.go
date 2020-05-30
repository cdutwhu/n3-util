package common

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// MustWriteFile :
func MustWriteFile(filename string, data []byte) {
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0700)
	}
	FailOnErr("%v", ioutil.WriteFile(filename, data, 0666))
}

// MustAppendFile :
func MustAppendFile(filename string, data []byte, newline bool) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		MustWriteFile(filename, []byte(""))
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	FailOnErr("%v", err)
	defer file.Close()

	if newline {
		data = append(append([]byte{}, '\n'), data...)
	}
	_, err = file.Write(data)
	FailOnErr("%v", err)
}
