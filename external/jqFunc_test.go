package external

import (
	"fmt"
	"os"
	"testing"
)

func TestJQ(t *testing.T) {
	bytes, err := os.ReadFile("../data/json/NAPCodeFrame.json")
	failOnErr("%v", err)
	fmtted := FmtJSONStr(string(bytes), "./utils/")
	// fmt.Println(fmtted)
	mustWriteFile("jq-fmt.json", []byte(fmtted))
	fmt.Println("OK")
}
