package csv

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	cmn "github.com/cdutwhu/json-util/common"
)

// ReadCSVFile : read the content of CSV File
func ReadCSVFile(path string, save bool, savePaths ...string) string {
	csvFile, err := os.Open(path)
	cmn.FailOnErr("The file is not found || wrong root : %v", err)
	defer csvFile.Close()
	bytes := ReadCSV(csvFile)
	if save {
		if len(savePaths) == 0 {
			newFileName := filepath.Base(path)
			newFileName = newFileName[0:len(newFileName)-len(filepath.Ext(newFileName))] + ".json"
			savepath := filepath.Join(filepath.Dir(path), newFileName)
			cmn.MustWriteFile(savepath, bytes)
		}
		for _, savepath := range savePaths {
			cmn.MustWriteFile(savepath, bytes)
		}
	}
	return string(bytes)
}

// ReadCSV to
func ReadCSV(r io.Reader) []byte {
	content, _ := csv.NewReader(r).ReadAll()
	cmn.FailOnErrWhen(len(content) < 1, "%v", fEf("Something wrong, the file maybe empty or length of the lines are not the same"))

	headersArr := make([]string, 0)
	for _, headE := range content[0] {
		headersArr = append(headersArr, headE)
	}

	//Remove the header row
	content = content[1:]

	// var buffer bytes.Buffer
	var buffer strings.Builder
	buffer.WriteString("[")
	for i, d := range content {
		buffer.WriteString("{")
		for j, y := range d {
			buffer.WriteString(`"` + headersArr[j] + `":`)
			_, fErr := strconv.ParseFloat(y, 32)
			_, bErr := strconv.ParseBool(y)
			if fErr == nil {
				buffer.WriteString(y)
			} else if bErr == nil {
				buffer.WriteString(strings.ToLower(y))
			} else {
				buffer.WriteString((`"` + y + `"`))
			}
			//end of property
			if j < len(d)-1 {
				buffer.WriteString(",")
			}
		}
		//end of object of the array
		buffer.WriteString("}")
		if i < len(content)-1 {
			buffer.WriteString(",")
		}
	}

	buffer.WriteString(`]`)
	rawMessage := json.RawMessage(buffer.String())
	jsonstr, err := json.MarshalIndent(rawMessage, "", "  ")
	cmn.FailOnErr("%v", err)
	return jsonstr
}
