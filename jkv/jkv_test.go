package jkv

import (
	"testing"
	"time"
)

func TestScan(t *testing.T) {
	defer trackTime(time.Now())
	json := fmtJSONFile("../data/json/NAPCodeFrame.json", 2)
	jkv := NewJKV(json, "", false)
	LVL, mLvlFParr, mFPosLvl, _ := jkv.scan(13)
	fPln("levels:", LVL)
	for k, v := range mLvlFParr {
		fPln(k, v)
	}
	for k, v := range mFPosLvl {
		fPln(k, v)
	}
}

func TestFieldByPos(t *testing.T) {
	defer trackTime(time.Now())
	json := fmtJSONFile("../data/json/NAPCodeFrame.json", 2)
	jkv := NewJKV(json, "", false)
	LVL, mLvlFParr, _, _ := jkv.scan(13)
	// for k, v := range mLvlFParr {
	// 	fPln(k, v)
	// }
	mFPosFNameList := jkv.fields(mLvlFParr)
	for i := 1; i <= LVL; i++ {
		fPln("---------------->", i)
		mFPosFName := mFPosFNameList[i]
		for k, v := range mFPosFName {
			_, t := jkv.fValueType(k)
			fPf("%-8d%-20s%-10s\n", k, v, t.Str())
			// if t.IsPrimitive() {
			// 	fPf("%-8d%-20s%-10s\n", k, v, t.Str())
			// } else {
			// 	fPf("%-8d%-20s\n", k, v)
			// }
		}
	}
}

func TestFType(t *testing.T) {
	defer trackTime(time.Now())
	json := fmtJSONFile("../data/json/NAPCodeFrame.json", 2)
	jkv := NewJKV(json, "", false)
	value, typ := jkv.fValueType(1617)
	fPln(typ.Str())
	if typ == ARR|OBJ {
		fPln(fValuesOnObjList(value)[1])
	}
}

func TestInit(t *testing.T) {
	defer trackTime(time.Now())
	json := fmtJSONFile("../data/json/NAPCodeFrame.json", 2)
	NewJKV(json, "", false)
	fPln("break")
}

func TestWrap(t *testing.T) {
	defer trackTime(time.Now())
	json := fmtJSONFile("../data/json/xapi1.json", 2)
	jkv := NewJKV(json, "root", false)
	fPln("--- Init ---")
	fPln(jkv.JSON)
}

func TestUnfold(t *testing.T) {
	defer trackTime(time.Now())

	json := fmtJSONFile("../data/json/xapi1.json", 2)
	jkv := NewJKV(json, "root", false)
	fPln("--- Init ---")
	fPln(jkv.Wrapped)
	fPln(jkv.Unfold(0, nil))

	// fPln(jkv.mOIDLvl["fe7262a928bbe05f8a42bab98ebec56a8e1e9379"])
	// fPln(jkv.mOIDLvl["887450b46a52ccad78f6a74f34c2699c649b17cd"]).

	fPln(" -------------------------------------- ")

	jkv = jkv.UnwrapDefault()
	// fPln(jkv.Unfold(0, nil))
	fPln(jkv.JSON)
}
