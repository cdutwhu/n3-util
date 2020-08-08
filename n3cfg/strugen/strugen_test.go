package strugen

import (
	"io/ioutil"
	"testing"
)

func TestScanToml(t *testing.T) {
	bytes, err := ioutil.ReadFile("../../_data/toml/test.toml")
	failP1OnErr("%v", err)
	toml := string(bytes)
	lines := splitLn(toml)

	attrs1, attrs2 := scanToml(lines)
	fPln(attrs1)
	fPln(attrs2)
}

func TestAttrsRange(t *testing.T) {
	bytes, err := ioutil.ReadFile("../../_data/toml/test.toml")
	failP1OnErr("%v", err)
	toml := string(bytes)
	lines := splitLn(toml)

	m := attrsRange(lines)
	for k, v := range m {
		fPln(k, v)
	}
}

func TestGenStruct(t *testing.T) {
	GenStruct("../../_data/toml/test.toml", "Config", "bank", "../bank/CfgSvr.go")
}

func TestAddConfig(t *testing.T) {
	file := AddCfg2Bank("qmiao", "../../_data/toml/test.toml", "config", "Svr")
	fPln(file)
}
