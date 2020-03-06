package jkv

import (
	"io/ioutil"
	"sync"
	"testing"
	"time"

	cmn "github.com/cdutwhu/json-util/common"
)

func TestJSONPolicy(t *testing.T) {
	defer cmn.TrackTime(time.Now())

	data := FmtJSONFile("../_data/NAPCodeFrame.json", 2)
	mask1 := FmtJSONFile("../_data/NAPCodeFrameMaskP.json", 2)
	mask2 := FmtJSONFile("../_data/NAPCodeFrameMaskPcopy.json", 2)

	cmn.FailOnErrWhen(data == "", "%v", fEf("input data is empty, check its path"))
	cmn.FailOnErrWhen(mask1 == "", "%v", fEf("input mask1 is empty, check its path"))
	cmn.FailOnErrWhen(mask2 == "", "%v", fEf("input mask2 is empty, check its path"))

	jkvM1 := NewJKV(mask1, "root", false)
	jkvM2 := NewJKV(mask2, "root", false)

	if MaybeJSONArr(data) {
		jsonArr := SplitJSONArr(data, 2)
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
		ioutil.WriteFile("array.json", []byte(MakeJSONArr(jsonList...)), 0666)

	} else {

		jkvD := NewJKV(data, "root", false)
		maskroot, _ := jkvD.Unfold(0, jkvM1)
		jkvMR := NewJKV(maskroot, "", false)
		jkvMR.Wrapped = jkvD.Wrapped
		json := jkvMR.UnwrapDefault().JSON
		json = FmtJSON(json, 2)

		jkvD = NewJKV(json, "root", false)
		maskroot, _ = jkvD.Unfold(0, jkvM2)
		jkvMR = NewJKV(maskroot, "", false)
		jkvMR.Wrapped = jkvD.Wrapped
		json = jkvMR.UnwrapDefault().JSON
		json = FmtJSON(json, 2)

		ioutil.WriteFile("single.json", []byte(json), 0666)
	}
}

func TestJSONPolicy1(t *testing.T) {
	defer cmn.TrackTime(time.Now())

	data := FmtJSONFile("../_data/test.json", 2)
	mask1 := FmtJSONFile("../_data/1.json", 2)

	cmn.FailOnErrWhen(data == "", "%v", fEf("input data is empty, check its path"))
	cmn.FailOnErrWhen(mask1 == "", "%v", fEf("input mask1 is empty, check its path"))

	jkvM1 := NewJKV(mask1, "root", false)

	jkvD := NewJKV(data, "root", false)
	maskroot, _ := jkvD.Unfold(0, jkvM1)
	jkvMR := NewJKV(maskroot, "", false)
	jkvMR.Wrapped = jkvD.Wrapped
	json := jkvMR.UnwrapDefault().JSON
	json = FmtJSON(json, 2)

	ioutil.WriteFile("single.json", []byte(json), 0666)
}
