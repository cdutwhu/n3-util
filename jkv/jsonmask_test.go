package jkv

import (
	"io/ioutil"
	"sync"
	"testing"
	"time"

	"github.com/cdutwhu/n3-util/n3err"
)

func TestJSONPolicy(t *testing.T) {
	defer trackTime(time.Now())

	data := fmtJSONFile("../_data/json/NAPCodeFrame.json", 2)
	mask1 := fmtJSONFile("../_data/json/NAPCodeFrameMaskP.json", 2)
	mask2 := fmtJSONFile("../_data/json/NAPCodeFrameMaskPcopy.json", 2)

	failOnErrWhen(data == "", "%v: input empty, check path", n3err.PARAM_INVALID)
	failOnErrWhen(mask1 == "", "%v: mask1 empty, check path", n3err.PARAM_INVALID)
	failOnErrWhen(mask2 == "", "%v: mask2 empty, check path", n3err.PARAM_INVALID)

	jkvM1 := NewJKV(mask1, "root", false)
	jkvM2 := NewJKV(mask2, "root", false)

	if maybeJSONArr(data) {
		jsonArr := splitJSONArr(data, 2)
		wg := sync.WaitGroup{}
		wg.Add(len(jsonArr))
		jsonList := make([]string, len(jsonArr))
		for i, json := range jsonArr {
			go func(i int, json string) {
				defer wg.Done()
				jkvD := NewJKV(json, "root", false)
				maskroot, _ := jkvD.Unfold(0, jkvM1)
				jkvMR := NewJKV(maskroot, "", false)
				jkvMR.Wrapped = jkvD.Wrapped
				jsonList[i] = jkvMR.UnwrapDefault().JSON
			}(i, json)
		}
		wg.Wait()
		ioutil.WriteFile("array.json", []byte(makeJSONArr(jsonList...)), 0666)

	} else {

		jkvD := NewJKV(data, "root", false)
		maskroot, _ := jkvD.Unfold(0, jkvM1)
		jkvMR := NewJKV(maskroot, "", false)
		jkvMR.Wrapped = jkvD.Wrapped
		json := jkvMR.UnwrapDefault().JSON
		json = fmtJSON(json, 2)

		jkvD = NewJKV(json, "root", false)
		maskroot, _ = jkvD.Unfold(0, jkvM2)
		jkvMR = NewJKV(maskroot, "", false)
		jkvMR.Wrapped = jkvD.Wrapped
		json = jkvMR.UnwrapDefault().JSON
		json = fmtJSON(json, 2)

		ioutil.WriteFile("single.json", []byte(json), 0666)
	}
}

func TestJSONPolicy1(t *testing.T) {
	defer trackTime(time.Now())

	data := fmtJSONFile("../_data/json/NAPCodeFrame.json", 2)
	mask := fmtJSONFile("../_data/json/NAPCodeFrameMask.json", 2)
	failOnErrWhen(data == "", "%v: data empty, check path", n3err.PARAM_INVALID)
	failOnErrWhen(mask == "", "%v: mask empty, check path", n3err.PARAM_INVALID)

	jkvD := NewJKV(data, "root", true) // data must wrap a level to do "simple" data & mask
	jkvM := NewJKV(mask, "", false)
	maskroot, _ := jkvD.Unfold(0, jkvM)
	jkvMR := NewJKV(maskroot, "", false)
	jkvMR.Wrapped = jkvD.Wrapped
	json := jkvMR.UnwrapDefault().JSON
	json = fmtJSON(json, 2)

	fPln(json)
	ioutil.WriteFile("single.json", []byte(json), 0666)
}
