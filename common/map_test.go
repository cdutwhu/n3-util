package common

import "testing"

func TestMapPrint(t *testing.T) {
	MapPrint(map[string]string{
		// "a": "b",
		"3": "4 a",
		"5": "b sss",
		"7": "B    sss",
		"1": "2  2345678  223",
	})
}
