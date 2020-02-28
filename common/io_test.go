package common

import (
	"testing"
)

func TestIO(t *testing.T) {
	// err := ioutil.WriteFile("./a/b/c.txt", []byte("hello"), 0666)
	// FailOnErr("%v", err)

	WriteFile("./a/b/c.txt", []byte("hello"), 0666)
}
