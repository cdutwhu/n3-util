package external

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestJQ(t *testing.T) {

	bytes, err := ioutil.ReadFile("../_data/json/NAPCodeFrame.json")
	failOnErr("%v", err)
	fmtted := FmtJSONStr(string(bytes), "./utils/")
	// fmt.Println(fmtted)
	ioutil.WriteFile("why.json", []byte(fmtted), 0666)
	fmt.Println("OK2")
}
