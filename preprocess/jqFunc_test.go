package preprocess

import (
	"fmt"
	"io/ioutil"
	"testing"

	cmn "github.com/cdutwhu/json-util/common"
)

// func MakeRTJQ() {
// 	switch runtime.GOOS {
// 	case "windows":
// 	case "linux":
// 	case "darwin":
// 	}
// }

func TestJQ(t *testing.T) {

	// fmt.Println(prepareJQ("../", "./", "./utils/"))
	// fmt.Println(os.Getwd())

	// fmt.Println(FmtJSONStr("{\"abc\": 123}", "../", "./", "./utils/"))

	// if data, err := ioutil.ReadFile("../data/sample.json"); err == nil {
	// 	// fmt.Println(string(data))
	// 	fmt.Println(FmtJSONStr(string(data), "../", "./", "./utils/"))
	// } else {
	// 	fmt.Println(err.Error())
	// }

	// formatted := FmtJSONFile("../../data/xapi.json", "../", "./", "../build/Linux64/")
	// ioutil.WriteFile("fmt.json", []byte(formatted), 0666)

	// formatted := FmtJSONFile("../_data/why.json", "./utils/")
	// ioutil.WriteFile("./whyfmt.json", []byte(formatted), 0666)
	// fmt.Println("OK1")

	bytes, err := ioutil.ReadFile("../_data/single quote.json")
	cmn.FailOnErr("%v", err)
	fmtted := FmtJSONStr(string(bytes), "./utils/")
	// fmt.Println(fmtted)
	ioutil.WriteFile("single quote.json", []byte(fmtted), 0666)
	fmt.Println("OK2")

	// FmtJSONFile("../data/xapi1.json", "../", "./", "./utils/")
}
