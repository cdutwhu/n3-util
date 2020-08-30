package strugen

import (
	"io/ioutil"
	"testing"
)

func TestScanToml(t *testing.T) {
	bytes, err := ioutil.ReadFile("../../data/toml/test.toml")
	failP1OnErr("%v", err)
	lines := splitLn(string(bytes))
	attrs1, attrs2 := scanToml(lines)
	fPln(attrs1)
	fPln(attrs2)
}

func TestAttrsRange(t *testing.T) {
	bytes, err := ioutil.ReadFile("../../data/toml/test.toml")
	failP1OnErr("%v", err)
	lines := splitLn(string(bytes))
	m := attrsRange(lines)
	for k, v := range m {
		fPln(k, v)
	}
}

func TestGenStruct(t *testing.T) {
	GenStruct("../../data/toml/test.toml", "Config", "n3cfg", "../Config.go")
}
