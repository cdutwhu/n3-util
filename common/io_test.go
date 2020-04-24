package common

import (
	"testing"
)

func TestMustWriteFile(t *testing.T) {
	MustWriteFile("./a/b/c.txt", []byte("hello"))
}

func TestMustAppendFile(t *testing.T) {
	MustAppendFile("./a/b/d.txt", []byte("hello"), true)
}
