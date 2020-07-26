package n3json

import (
	"testing"
	"time"

	"github.com/cdutwhu/n3-util/n3err"
)

func TestInnerFmt(t *testing.T) {
	s := `{
	      a,
	      b
	    }`

	s = InnerFmt(s)
	fPln(s)
}

func TestSplitArr(t *testing.T) {
	defer trackTime(time.Now())

	// jArrStr := FmtFile("../_data/xapi.json", 2)
	// jArrStr := FmtFile("../../Server/config/meta.json", 2)
	// failOnErrWhen(jArrStr == "", "%v", n3err.FILE_EMPTY)

	jArrStr := FmtFile("../_data/json/xapi.json", 2)

	if arr := SplitArr(jArrStr, 2); arr != nil {
		jMergedStr := MakeArr(arr...)
		// fPln(jMergedStr)
		failOnErrWhen(jArrStr != jMergedStr, "%v: MakeArr", n3err.INTERNAL)
	} else {
		failOnErr("%v", n3err.JSON_ARRAY_NOT_FMT)
	}
}

func TestScalarSel(t *testing.T) {
	json := FmtFile("../_data/csv/Questions.json", 2)
	result := ScalarSelX(json, "module_version_id", "question_id", "category", "display_order", "question_type", "actual_answer")
	fPln(result)
}

func BenchmarkScalarSelX(b *testing.B) {
	json := FmtFile("../_data/csv/data1.json", 2)
	for n := 0; n < b.N; n++ {
		ScalarSelX(json, "Id", "Name", "Age1")
	}
}

// -------------------------- //

func TestJoin(t *testing.T) {
	jastrL := FmtFile("../_data/csv/Modules.json", 2)
	jastrR := FmtFile("../_data/csv/Substrands.json", 2)
	result, pairs := ArrJoin(jastrL, "substrand_id", jastrR, "substrand_id", "")
	for _, pair := range pairs {
		fPln(pair, "joined")
	}
	mustWriteFile("../_data/csv/Modules-Substrands.json", []byte(result))
}
