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
	ioutil.WriteFile("jq-fmt.json", []byte(fmtted), 0666)
	fmt.Println("OK")
}
