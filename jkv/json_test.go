package jkv

import (
	"testing"
	"time"

	cmn "github.com/cdutwhu/json-util/common"
)

func TestSplitJSONArr(t *testing.T) {
	defer cmn.TrackTime(time.Now())

	// jArrStr := FmtJSONFile("../_data/xapi.json", 2)
	// jArrStr := FmtJSONFile("../../Server/config/meta.json", 2)
	// cmn.FailOnErrWhen(jArrStr == "", "%v", fEf("Read JSON file error"))

	jArrStr := FmtJSONFile("../_data/xapi.json", 2)

	if arr := SplitJSONArr(jArrStr, 2); arr != nil {
		jMergedStr := MakeJSONArr(arr...)
		// fPln(jMergedStr)
		cmn.FailOnErrWhen(jArrStr != jMergedStr, "%v", fEf("MakeJSONArr Error"))
	} else {
		cmn.FailOnErr("%v", fEf("non-formatted json array"))
	}
}

func TestJSONScalarSel(t *testing.T) {
	json := FmtJSONFile("../csv/data/Questions.json", 2)
	result := JSONScalarSelX(json, "module_version_id", "question_id", "category", "display_order", "question_type", "actual_answer")
	fPln(result)
}

func BenchmarkJSONScalarSelX(b *testing.B) {
	json := FmtJSONFile("../csv/data1.json", 2)
	for n := 0; n < b.N; n++ {
		JSONScalarSelX(json, "Id", "Name", "Age1")
	}
}
