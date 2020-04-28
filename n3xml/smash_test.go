package n3xml

import (
	"io/ioutil"
	"testing"
)

func TestSmash(t *testing.T) {
	bytes, err := ioutil.ReadFile("../_data/xml/siftest.xml")
	failOnErr("%v", err)
	SmashSave(string(bytes), "./sif")
}
