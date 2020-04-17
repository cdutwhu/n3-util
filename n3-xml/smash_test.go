package xml

import (
	"io/ioutil"
	"testing"

	cmn "github.com/cdutwhu/json-util/common"
)

func TestSmash(t *testing.T) {
	bytes, err := ioutil.ReadFile("./siftest.xml")
	cmn.FailOnErr("%v", err)
	SmashSave(string(bytes), "./sif")
}
