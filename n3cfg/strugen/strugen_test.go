package strugen

import (
	"os"
	"testing"
)

var (
	toml  = "../../data/toml/test.toml"
	toml1 = "../../data/toml/test1.toml"
)

func TestScanToml(t *testing.T) {
	bytes, err := os.ReadFile(toml)
	failP1OnErr("%v", err)
	lines := splitLn(string(bytes))
	attrs1, attrs2 := scanToml(lines)
	fPln(attrs1)
	fPln(attrs2)
}

func TestAttrsRange(t *testing.T) {
	bytes, err := os.ReadFile(toml)
	failP1OnErr("%v", err)
	lines := splitLn(string(bytes))
	for k, v := range attrsRange(lines) {
		fPln(k, v)
	}
}

func TestAttrTypes(t *testing.T) {
	bytes, err := os.ReadFile(toml)
	failP1OnErr("%v", err)
	lines := splitLn(string(bytes))
	for k, v := range attrTypes(lines, "") {
		fPln(k, v)
	}
}

func TestGenStruct(t *testing.T) {
	enableWarnDetail(false)
	cfgsrc := "../Config.go"
	os.Remove(cfgsrc)
	GenStruct(toml, "Config", "n3cfg", cfgsrc)
	GenStruct(toml1, "Config1", "n3cfg", cfgsrc)
	GenNewCfg(cfgsrc)
}
