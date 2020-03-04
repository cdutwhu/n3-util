package csv

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestCSV2JSON(t *testing.T) {
	path := flag.String("path", "./data.csv", "Path of the file")
	flag.Parse()
	ReadCSVFile(*path, true, "data1.json")
	fmt.Println(strings.Repeat("=", 10), "Done", strings.Repeat("=", 10))
}

func BenchmarkCSV2JSON(b *testing.B) {
	path := "./data.csv"
	for n := 0; n < b.N; n++ {
		csv, _ := os.Open(path)
		ReadCSV(csv)
		csv.Close()
	}
}
