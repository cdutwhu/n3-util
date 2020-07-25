package n3toml

import (
	"io/ioutil"
	"testing"
)

func TestScanToml(t *testing.T) {
	bytes, err := ioutil.ReadFile("../_data/toml/test.toml")
	failP1OnErr("%v", err)
	toml := string(bytes)
	lines := sSplit(toml, "\n")

	attrs1, attrs2 := scanToml(lines)
	fPln(attrs1)
	fPln(attrs2)
}

func TestAttrsRange(t *testing.T) {
	bytes, err := ioutil.ReadFile("../_data/toml/test.toml")
	failP1OnErr("%v", err)
	toml := string(bytes)
	lines := sSplit(toml, "\n")

	m := attrsRange(lines)
	for k, v := range m {
		fPln(k, v)
	}
}

func TestAttrTypes(t *testing.T) {
	TomlGenStruct("../_data/toml/test.toml", "OOO", "ppp", "../_data/toml/test111.go")
	TomlGenStruct("../_data/toml/test.toml", "", "", "")
}
