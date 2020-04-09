package jkv

import (
	"testing"

	cmn "github.com/cdutwhu/json-util/common"
)

func TestNewTraits(t *testing.T) {

	// if Trait4Scan != TraitScan {
	// 	panic("error 0")
	// }

	for i := range cmn.N(13) {
		fPln(i)
	}

	// for i, s := range StartOfObjArr(cmn.Iter2Slc(13)...) {
	// 	cmn.FailOnErrWhen(s != sTAOStart[i], "%v", fEf("err 1"))
	// }

	// for i, s := range EndOfObjArr(cmn.Iter2Slc(13)...) {
	// 	cmn.FailOnErrWhen(s != sTAOEnd[i], "%v", fEf("err 2"))
	// }
}
