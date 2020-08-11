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
	GenStruct("../../_data/toml/test.toml", "Config", "n3cfg", "../Config.go")
}

// echo 'password' | sudo -S env "PATH=$PATH" go test -v -count=1 ./ -run TestRegisterCfg
func TestRegisterCfg(t *testing.T) {
	prj, pkg := "n3-util", "Server"
	ok, file := RegisterCfg("qmiao", "../../_data/toml/test.toml", prj, pkg)
	fPln(ok, file)
}
