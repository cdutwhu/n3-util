package csv

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	cmn "github.com/cdutwhu/json-util/common"
)

func TestCSV2JSON(t *testing.T) {

	dir := "./data/"
	files, err := ioutil.ReadDir(dir)
	cmn.FailOnErr("%v", err)

	for _, file := range files {
		fName := dir + file.Name()
		fPln(fName)
		if !strings.HasSuffix(file.Name(), ".csv") {
			continue
		}
		File2JSON(fName, true, sReplaceAll(fName, ".csv", ".json"))
	}

	// path := flag.String("path", "./data/ModulePrerequisites.csv", "Path of the file")
	// flag.Parse()
	// File2JSON(*path, true, "data.json")
	// fmt.Println(strings.Repeat("=", 10), "Done", strings.Repeat("=", 10))
}

func BenchmarkCSV2JSON(b *testing.B) {
	path := "./data/data.csv"
	for n := 0; n < b.N; n++ {
		csv, _ := os.Open(path)
		Reader2JSON(csv)
		csv.Close()
	}
}
