package n3xml

import (
	"io/ioutil"
	"testing"
)

func TestSmash(t *testing.T) {
	bytes, err := ioutil.ReadFile("../data/xml/siftest346.xml")
	failOnErr("%v", err)
	SmashSave(string(bytes), "./sif346")
}
