package external

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestJQ(t *testing.T) {
	bytes, err := ioutil.ReadFile("../data/json/NAPCodeFrame.json")
	failOnErr("%v", err)
	fmtted := FmtJSONStr(string(bytes), "./utils/")
	// fmt.Println(fmtted)
	mustWriteFile("jq-fmt.json", []byte(fmtted))
	fmt.Println("OK")
}
