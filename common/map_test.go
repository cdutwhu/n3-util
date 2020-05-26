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

func TestMapFromStruct(t *testing.T) {
	s := struct {
		A string
		B int
		C bool
	}{A: "aa", B: 22, C: false}
	m := MapFromStruct(s)
	fPln(m)

	ks := MapKeys(m)
	fPln(ks)

	fPln(" ------------------- ")

	s1 := struct {
		A1 string
		B1 string
		C1 string
	}{A1: "aa", B1: "22", C1: "false"}
	m = MapFromStruct(s1)
	fPln(m)

	ks, vs := MapKVs(m)
	fPln(ks)
	fPln(vs)
}
