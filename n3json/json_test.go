package n3json

import (
	"testing"
	"time"
)

func TestInnerFmt(t *testing.T) {
	s := `{
	      a,
	      b
	    }`

	s, did := InnerFmt(s)
	fPln(s, did)
}

func TestSplitJSONArr(t *testing.T) {
	defer trackTime(time.Now())

	// jArrStr := FmtJSONFile("../_data/xapi.json", 2)
	// jArrStr := FmtJSONFile("../../Server/config/meta.json", 2)
	// failOnErrWhen(jArrStr == "", "%v", fEf("Read JSON file error"))

	jArrStr := FmtFile("../_data/xapi.json", 2)

	if arr := SplitArr(jArrStr, 2); arr != nil {
		jMergedStr := MakeArr(arr...)
		// fPln(jMergedStr)
		failOnErrWhen(jArrStr != jMergedStr, "%v", fEf("MakeJSONArr Error"))
	} else {
		failOnErr("%v", fEf("non-formatted json array"))
	}
}

func TestJSONScalarSel(t *testing.T) {
	json := FmtFile("../n3csv/data/Questions.json", 2)
	result := ScalarSelX(json, "module_version_id", "question_id", "category", "display_order", "question_type", "actual_answer")
	fPln(result)
}

func BenchmarkJSONScalarSelX(b *testing.B) {
	json := FmtFile("../n3csv/data1.json", 2)
	for n := 0; n < b.N; n++ {
		ScalarSelX(json, "Id", "Name", "Age1")
	}
}

// -------------------------- //

func TestJSONJoin(t *testing.T) {
	jastrL := FmtFile("../n3csv/data/Modules.json", 2)
	jastrR := FmtFile("../n3csv/data/Substrands.json", 2)
	result, pairs := ArrJoin(jastrL, "substrand_id", jastrR, "substrand_id", "")
	for _, pair := range pairs {
		fPln(pair, "joined")
	}
	mustWriteFile("../n3csv/data/Modules-Substrands.json", []byte(result))
}
