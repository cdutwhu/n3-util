package n3csv

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// File2JSON : read the content of CSV File
func File2JSON(path string, vertical, save bool, savePaths ...string) (string, []string) {
	csvFile, err := os.Open(path)
	failOnErr("The file is not found || wrong root : %v", err)
	defer csvFile.Close()
	jsonstr, headers := Reader2JSON(csvFile)

	if vertical {
		jsonstr = JSONScalarSelX(jsonstr, headers...)
	}

	if save {
		if len(savePaths) == 0 {
			newFileName := filepath.Base(path)
			newFileName = newFileName[0:len(newFileName)-len(filepath.Ext(newFileName))] + ".json"
			savepath := filepath.Join(filepath.Dir(path), newFileName)
			mustWriteFile(savepath, []byte(jsonstr))
		}
		for _, savepath := range savePaths {
			mustWriteFile(savepath, []byte(jsonstr))
		}
	}
	return jsonstr, headers
}

// Reader2JSON to
func Reader2JSON(r io.Reader) (string, []string) {
	content, _ := csv.NewReader(r).ReadAll()
	failOnErrWhen(len(content) < 1, "%v", fEf("Failed, the file may be empty or length of the lines are not the same"))

	headersArr := make([]string, 0)
	for _, headE := range content[0] {
		headersArr = append(headersArr, headE)
	}

	//Remove the header row
	content = content[1:]

	// Set Column Type
	mColType := make(map[int]string)
	for _, d := range content {
		for j, y := range d {
			_, fErr := strconv.ParseFloat(y, 32)
			_, bErr := strconv.ParseBool(y)
			switch {
			case fErr == nil:
				mColType[j] = "numeric"
			case bErr == nil:
				mColType[j] = "bool"
			default:
				mColType[j] = "string"
			}
		}
	}
	//

	// var buffer bytes.Buffer
	var buffer strings.Builder
	buffer.WriteString("[")
	for i, d := range content {
		buffer.WriteString("{")
		for j, y := range d {
			buffer.WriteString(`"` + headersArr[j] + `":`)

			// _, fErr := strconv.ParseFloat(y, 32)
			// _, bErr := strconv.ParseBool(y)
			// if fErr == nil {
			// 	buffer.WriteString(y)
			// } else if bErr == nil {
			// 	buffer.WriteString(strings.ToLower(y))
			// } else {
			// 	buffer.WriteString((`"` + y + `"`))
			// }

			switch mColType[j] {
			case "numeric":
				buffer.WriteString(y)
			case "bool":
				buffer.WriteString(strings.ToLower(y))
			case "string":
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
	failOnErr("%v", err)
	return string(jsonstr), headersArr
}
