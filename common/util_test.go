package common

import (
	"testing"
)

func TestColor(t *testing.T) {
	fPln("\033[31mRed")
	fPln("\033[32mGreen")
	fPln("\033[34mBlue")
}
