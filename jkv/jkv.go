// ********** ALL are Based On JQ Formatted JSON ********** //

package jkv

import (
	"math"

	eg "github.com/cdutwhu/n3-util/n3errs"
)

// NewJKV :
func NewJKV(jsonstr, defroot string, mustWrap bool) *JKV {
	jkv := &JKV{
		JSON: jsonstr,
		LsL12Fields: [][]string{
			{}, {}, {},
		},
		lsLvlIPaths: [][]string{
			{}, {}, {}, {}, {},
			{}, {}, {}, {}, {},
			{}, {}, {}, {}, {},
			{}, {}, {}, {}, {},
			{}, {}, {}, {}, {},
		},
		mPathMAXIdx:   make(map[string]int),      //
		mIPathPos:     make(map[string]int),      //
		MapIPathValue: make(map[string]string),   //
		mIPathOID:     make(map[string]string),   //
		mOIDiPath:     make(map[string]string),   //
		mOIDObj:       make(map[string]string),   //
		mOIDLvl:       make(map[string]int),      // from 1 ...
		mOIDType:      make(map[string]JSONTYPE), // oid-type is OBJ or ARR|OBJ
	}
	jkv.init()
	if defroot == "" {
		return jkv
	}
	return jkv.wrapDefault(defroot, mustWrap)
}

// **************************************************************** //

// isJSON :
func (jkv *JKV) isJSON() bool {
	return isJSON(jkv.JSON)
}

// scan :                                 L   posarr     pos L
func (jkv *JKV) scan(depth int) (int, map[int][]int, map[int]int, error) {
	Lm, offset := -1, 0
	if s := jkv.JSON; jkv.isJSON() {
		mLvlFParr := make(map[int][]int)
		for i := 0; i <= LvlMax; i++ {
			mLvlFParr[i] = []int{}
		}
		mFPosLvl := make(map[int]int)

		indices := iter2Slc(depth)
		sTAOStart := StartOfObjArr(indices...)
		sTAOEnd := EndOfObjArr(indices...)

		// L0 : object
		if s[0] == '{' {
		NEXT:
			for i := 0; i < len(s); i++ {
				// modify levels for array-object
				if ok, _ := hasAnyPrefix(s[i:], sTAOStart...); ok {
					offset++
				}
				if ok, _ := hasAnyPrefix(s[i:], sTAOEnd...); ok {
					offset--
				}

				for j := 3; j <= 39; j += 2 {
					T, L := TL(j)
					e := i + j

					if e < len(s) && s[i:e] == T && s[e] == '"' { // xIn(s[e], StartTrait) {
						// remove fakes (remove string array)
						for k := e + 1; k < len(s)-1; k++ {
							if s[k] == '"' {
								if s[k+1] != ':' {
									continue NEXT
								}
								break
							}
						}

						L -= offset
						mLvlFParr[L] = append(mLvlFParr[L], e) // store '"' position
						mFPosLvl[e] = L
						continue NEXT
					}
				}
			}
		}

		// remove empty levels
		for i := LvlMax; i >= 0; i-- {
			if v := mLvlFParr[i]; len(v) == 0 {
				delete(mLvlFParr, i)
				continue
			}
			Lm = i
			break
		}

		return Lm, mLvlFParr, mFPosLvl, nil
	}
	return Lm, nil, nil, eg.JSON_INVALID
}

// fields :
func (jkv *JKV) fields(mLvlFPos map[int][]int) []map[int]string {
	s := jkv.JSON
	Ikeys, err := mapKeys(mLvlFPos)
	failOnErr("%v", err)
	keys := Ikeys.([]int)

	nLVL := keys[len(keys)-1]
	mFPosFNameList := []map[int]string{{}} // L0 is empty
	for L := 1; L <= nLVL; L++ {           // from L1 to Ln
		mFPosFName := make(map[int]string)
		for _, p := range mLvlFPos[L] {
			pe := p + 1
			for i := p + 1; s[i] != DQ; i++ {
				pe = i
			}
			mFPosFName[p] = s[p+1 : pe+1]
		}
		mFPosFNameList = append(mFPosFNameList, mFPosFName)
	}
	return mFPosFNameList
}

// pl2 -> pl1. pl1, pl2 are sorted.
func merge2fields(mFPosFName1, mFPosFName2 map[int]string) map[int]string {
	pl2Parent, pl2Path, iPos := make(map[int]string), make(map[int]string), 0

	Ipl1, err := mapKeys(mFPosFName1)
	failOnErr("%v", err)
	pl1 := Ipl1.([]int)

	Ipl2, err := mapKeys(mFPosFName2)
	failOnErr("%v", err)
	pl2 := Ipl2.([]int)

	for _, p2 := range pl2 {
		for i := iPos; i < len(pl1)-1; i++ {
			if p2 > pl1[i] && p2 < pl1[i+1] {
				iPos = i
				pl2Parent[p2] = mFPosFName1[pl1[i]]
				break
			}
		}
		if p2 > pl1[len(pl1)-1] {
			pl2Parent[p2] = mFPosFName1[pl1[len(pl1)-1]]
		}
		pl2Path[p2] = pl2Parent[p2] + pLinker + mFPosFName2[p2]
	}

	Imap, err := mapsJoin(mFPosFName1, pl2Path)
	failOnErr("%v", err)
	return Imap.(map[int]string)
}

// rely on "fields outcome"
func fPaths(mFPosFNameList ...map[int]string) map[int]string {
	mFPosFPath := make(map[int]string)
	nL := len(mFPosFNameList)
	posLists := make([][]int, nL)
	for i, mFPosFName := range mFPosFNameList {
		if len(mFPosFName) == 0 {
			continue
		}

		IposList, err := mapKeys(mFPosFName)
		failOnErr("%v", err)
		posList := IposList.([]int)
		posLists[i] = posList
	}
	mFPosFNameMerge := mFPosFNameList[1]
	for i := 1; i < nL-1; i++ {
		mFPosFNameMerge = merge2fields(mFPosFNameMerge, mFPosFNameList[i+1])
		mFPosFPath = mFPosFNameMerge
	}
	return mFPosFPath
}

// ********************************************************** //

// fValuesOnObjList :
func fValuesOnObjList(strObjlist string) (objlist []string) {
	L, mLPStart, mLPEnd := 0, make(map[int][]int), make(map[int][]int)
	for p := 0; p < len(strObjlist); p++ {
		c := strObjlist[p]
		if c == '{' {
			L++
			mLPStart[L] = append(mLPStart[L], p)
		}
		if c == '}' {
			mLPEnd[L] = append(mLPEnd[L], p)
			L--
		}
	}
	pstarts, pends := mLPStart[1], mLPEnd[1]
	for i := 0; i < len(pstarts); i++ {
		s, e := pstarts[i], pends[i]
		objlist = append(objlist, strObjlist[s:e+1])
	}
	return objlist
}

// fValueType :
func (jkv *JKV) fValueType(p int) (v string, t JSONTYPE) {
	getV := func(str string, s int) string {
		for i := s + 1; i < len(str); i++ {
			if ok, _ := hasAnyPrefix(str[i:], Trait1EndV, Trait2EndV); ok {
				return str[s:i]
			}
		}
		panic("Shouldn't be here @ getV")
	}
	getOV := func(str string, s int) string {
		nLCB, nRCB := 0, 0
		for i := s; i < len(str); i++ {
			switch str[i] {
			case '{':
				nLCB++
			case '}':
				nRCB++
			}
			ok, err := hasAnyPrefix(str[i:], "},\n", "}\n")
			failOnErr("%v", err)
			if nLCB == nRCB && ok {
				return str[s : i+1]
			}
		}
		panic("Shouldn't be here @ getOV")
	}
	getAV := func(str string, s int) string {
		nLBB, nRBB := 0, 0
		for i := s; i < len(str); i++ {
			switch str[i] {
			case '[':
				nLBB++
			case ']':
				nRBB++
			}

			ok, err := hasAnyPrefix(str[i:], "],\n", "]\n")
			failOnErr("%v", err)
			if nLBB == nRBB && ok {
				return str[s : i+1]
			}
		}
		panic("Shouldn't be here @ getAV")
	}

	s := jkv.JSON
	v1c, pv := byte(0), 0
	for i := p; i < len(s); i++ {
		if sHasPrefix(s[i:], TraitFV) {
			pv = i + len(TraitFV)
			v1c = s[pv]
			break
		}
	}
	switch v1c {
	case DQ:
		t, v = STR, getV(s, pv)
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
		t, v = NUM, getV(s, pv)
	case 't', 'f':
		t, v = BOOL, getV(s, pv)
	case 'n':
		t, v = NULL, getV(s, pv)
	case '{':
		t, v = OBJ, getOV(s, pv)
	case '[':
		t, v = ARR, getAV(s, pv)
		{
			for i := pv + 1; i < len(s); i++ {
				c := s[i]
				if c == LF || c == SP {
					continue
				}
				switch c {
				case DQ:
					t |= STR
				case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
					t |= NUM
				case 't', 'f':
					t |= BOOL
				case 'n':
					t |= NULL
				case '{':
					t |= OBJ
				default:
					panic("Invalid JSON array element type")
				}
				break
			}
		}
	default:
		panic(fSf("[%d] @ Invalid JSON element type", p))
	}
	return
}

// pathType :
func (jkv *JKV) pathType(fPath string, psSort []int, mFPosFPath map[int]string) JSONTYPE {
	for _, p := range psSort {
		if fPath == mFPosFPath[p] {
			_, t := jkv.fValueType(p)
			return t
		}
	}
	panic("Shouldn't be here @ pathType")
}

// init : prepare <>
func (jkv *JKV) init() error {
	const scanDepth = 13
	if _, mLvlFParr, _, err := jkv.scan(scanDepth); err == nil {
		lsMapFPosFName := jkv.fields(mLvlFParr)

		for iL, mPN := range lsMapFPosFName {
			// fPln("<------Level------>", iL)
			for _, name := range mPN {
				// ----- //
				// v, t := jkv.fValueType(p)
				// if !t.IsLeafValue() {
				// 	oid := uuid.New().String()
				// 	v = oid
				// }
				// fPln(t.Str(), name, v)
				// ----- //

				if iL < len(jkv.LsL12Fields) {
					jkv.LsL12Fields[iL] = append(jkv.LsL12Fields[iL], name)
				}
			}
		}

		mFPath := fPaths(lsMapFPosFName...)
		if len(mFPath) == 0 {
			return err
		}

		Ikeys, err := mapKeys(mFPath)
		failOnErr("%v", err)
		for _, p := range Ikeys.([]int) {
			v, t := jkv.fValueType(p)

			oid := ""
			if !t.IsLeafValue() {
				failOnErrWhen(!isJSON(v), "%v: fetching value error", eg.INTERNAL)
				oid = hash(v)
				jkv.mOIDObj[oid] = v
				v = oid
				if t.IsObj() || t.IsObjArr() {
					jkv.mOIDType[oid] = t
				}
			}

			fp := mFPath[p]
			fip := fSf("%s@%d", fp, jkv.mPathMAXIdx[fp])
			jkv.mPathMAXIdx[fp]++
			jkv.MapIPathValue[fip] = v
			jkv.mIPathPos[fip] = p
			// fPf("DEBUG: %-5d%-5d[%-7s]  [%-60s]  %s\n", i, p, t.Str(), fip, v)

			if !t.IsLeafValue() {
				jkv.mIPathOID[fip] = oid
				jkv.mOIDiPath[oid] = fip
			}
		}

		//
		for iPath := range jkv.mIPathOID {
			n := sCount(iPath, pLinker) + 1
			jkv.lsLvlIPaths[n] = append(jkv.lsLvlIPaths[n], iPath)
			// fPf("%s [%d] %s\n", oid, n, iPath)
		}

		for i := 1; i < len(jkv.lsLvlIPaths); i++ {
			if Ls, LsPrev := jkv.lsLvlIPaths[i], jkv.lsLvlIPaths[i-1]; len(Ls) > 0 && len(LsPrev) > 0 {
				for _, iPathP := range LsPrev {
					pathP := rmTailFromLast(iPathP, "@")
					chk := pathP + pLinker
					for _, iPath := range Ls {
						if sHasPrefix(iPath, chk) {
							oidP, oid := jkv.mIPathOID[iPathP], jkv.mIPathOID[iPath]
							objP, obj := jkv.mOIDObj[oidP], jkv.mOIDObj[oid]
							jkv.mOIDObj[oidP] = sReplaceAll(objP, obj, oid)
							jkv.mOIDLvl[oidP], jkv.mOIDLvl[oid] = i-1, i
						}
					}
				}
			}
		}

		// [obj-arr whole value string] -> [aoID arr string]
		for oid := range jkv.mOIDObj {
			if strOIDlist := jkv.aoID2oIDlist(oid); strOIDlist != "" {
				jkv.mOIDObj[oid] = strOIDlist
				lvl := jkv.mOIDLvl[oid]
				aoIDs, err := oIDlistStr2oIDlist(strOIDlist)
				failOnErr("%v", err)
				for _, aoID := range aoIDs {
					jkv.mOIDLvl[aoID] = lvl
				}
			}
		}

		return nil
	}

	return eg.INTERNAL_SCAN_ERR
}

// aoID2oIDlist : only can be used after mOIDType assigned
func (jkv *JKV) aoID2oIDlist(aoID string) string {
	if typ, ok := jkv.mOIDType[aoID]; ok && typ.IsObjArr() {
		strObjlist := jkv.mOIDObj[aoID]
		objlist := fValuesOnObjList(strObjlist)
		for _, obj := range objlist {
			oid := hash(obj)
			jkv.mOIDType[oid] = OBJ
			jkv.mOIDiPath[oid] = jkv.mOIDiPath[aoID]
			jkv.mOIDLvl[oid] = jkv.mOIDLvl[aoID]
			jkv.mOIDObj[oid] = obj
			strObjlist = sReplace(strObjlist, obj, oid, 1)
		}
		return strObjlist
	}
	return ""
}

// oIDlistStr2oIDlist : string: "[ ****, ****, **** ]" => [ ****, ****, **** ]
func oIDlistStr2oIDlist(aoIDStr string) (oidlist []string, err error) {
	nComma := sCount(aoIDStr, ",")
	oidlist = hashRExp.FindAllString(aoIDStr, -1)
	if aoIDStr[0] != '[' || aoIDStr[len(aoIDStr)-1] != ']' || (oidlist != nil && len(oidlist) != nComma+1) {
		return nil, eg.PARAM_INVALID_FMT
	}
	return oidlist, nil
}

// ******************************************** //

// wrapDefault :
func (jkv *JKV) wrapDefault(root string, must bool) *JKV {
	if len(jkv.LsL12Fields[1]) == 1 && !must {
		return jkv
	}
	json := jkv.JSON
	if !sHasSuffix(json, "\n") {
		json += "\n"
	}

	jsonInd, _ := indent(json, 2, true)
	rooted1 := fSf("{\n  \"%s\": %s}\n", root, jsonInd)
	rooted2 := fSf("{\n  \"%s\": %s}\n", root, json)
	rooted2 = fmtJSON(rooted2, 2) + "\n"
	if rooted1 != rooted2 {
		mustWriteFile("./root1.json", []byte(rooted1))
		mustWriteFile("./root2.json", []byte(rooted2))
	}
	failOnErrWhen(rooted1 != rooted2, "%v: rooted", eg.INTERNAL)

	// fPln(" ----------------------------------------------- ")
	jkvR := NewJKV(rooted1, "", must)
	jkvR.Wrapped = true
	return jkvR
}

// UnwrapDefault :
func (jkv *JKV) UnwrapDefault() *JKV {
	if !jkv.Wrapped {
		return jkv
	}
	json := jkv.JSON
	i, j, n1, n2 := 0, len(json)-1, 0, 0
	for i, j = 0, len(json)-1; i < len(json) && j >= 0; {
		if n1 < 2 {
			if json[i] == '{' {
				n1++
			}
			i++
		}
		if n2 < 2 {
			if json[j] == '}' {
				n2++
			}
			j--
		}
		if n1 == 2 && n2 == 2 {
			break
		}
	}

	unRooted1, _ := fmtInnerJSON(json[i-1 : j+2])
	unRooted1 += "\n"
	// fPln(unRooted1)
	unRooted2 := fmtJSON(json[i-1:j+2], 2)
	unRooted2 += "\n"
	// fPln(unRooted2)
	if unRooted1 != unRooted2 {
		mustWriteFile("./unroot1.json", []byte(unRooted1))
		mustWriteFile("./unroot2.json", []byte(unRooted2))
	}
	failOnErrWhen(unRooted1 != unRooted2, "%v: unRooted", eg.INTERNAL)

	jkvUnR := NewJKV(unRooted1, "", false)
	jkvUnR.Wrapped = false
	return jkvUnR
}

// Unfold :
func (jkv *JKV) Unfold(toLvl int, mask *JKV) (string, int) {

	frame := ""
	if len(jkv.lsLvlIPaths[1]) == 0 {
		frame = ""
	} else if len(jkv.lsLvlIPaths[1]) != 0 && len(jkv.lsLvlIPaths[2]) == 0 {
		frame = jkv.JSON
	} else {
		firstField := jkv.lsLvlIPaths[1][0]
		lvl1path := rmTailFromLast(firstField, "@")
		oid := jkv.MapIPathValue[firstField]
		frame = fSf("{\n  \"%s\": %s\n}", lvl1path, oid)
	}

	//	maskLvlFields := projectV(MapKeys(mask.MapIPathValue).([]string), pLinker, "", "@")

	// expanding ...
	iExp := 0
	for {
		iExp++

		// [object array whole oid] => [ oid, oid, oid ... ]
		for _, oid := range hashRExp.FindAllString(frame, -1) {
			if jkv.mOIDType[oid].IsObjArr() {
				frame = sReplaceAll(frame, oid, jkv.mOIDObj[oid])
			}
		}
		if toLvl == 1 && iExp == toLvl {
			return frame, iExp // DEBUG testing
		}

		if oIDlist := hashRExp.FindAllString(frame, -1); oIDlist != nil {
			for _, oid := range oIDlist {
				ss := sSplit(jkv.mOIDiPath[oid], pLinker)
				name := sSplit(ss[len(ss)-1], "@")[0]
				obj := jkv.mOIDObj[oid]
				objMasked := Mask(name, obj, mask)
				frame = sReplaceAll(frame, oid, objMasked)

				// [object array whole oid] => [ oid, oid, oid ... ]
				for _, oid := range hashRExp.FindAllString(obj, -1) {
					if jkv.mOIDType[oid].IsObjArr() {
						frame = sReplaceAll(frame, oid, jkv.mOIDObj[oid])
					}
				}
			}
			if toLvl > 1 && iExp+1 == toLvl {
				return frame, toLvl // DEBUG testing
			}

		} else {
			break
		}
	}

	failOnErrWhen(!isJSON(frame), "%v: UNFOLD ERROR", eg.INTERNAL)
	return frame, iExp
}

// Mask :
func Mask(name, obj string, mask *JKV) string {
	if mask == nil {
		return obj
	}

	// check current mask path is valid for current objTmp fields, P1/2
	objTmp, _ := fmtInnerJSON(obj)
	jkvTmp := NewJKV(objTmp, name, true)
	pathlistTmp := func(name, linker string, fields []string) (pathlist []string) {
		for _, f := range fields {
			pathlist = append(pathlist, name+linker+f)
		}
		return
	}(name, pLinker, jkvTmp.LsL12Fields[2])
	// END -- P1/2 //

	for path, valMask := range mask.MapIPathValue {
		path = rmTailFromLast(path, "@")

		// check current mask path is valid for current objTmp fields,
		// if AT LEAST ONE mask path is valid, let this path go through and make effect. P2/2
		flag := false
		for _, pathTmp := range pathlistTmp {
			if path != pathTmp && !sHasSuffix(path, pLinker+pathTmp) {
				continue
			}
			flag = true
			break
		}
		if !flag {
			continue
		}
		// END -- P2/2 //

		field := rmHeadToLast(path, pLinker)
		lookfor := fSf("\"%s%s", field, TraitFV)

		if i := sIndex(obj, lookfor); i > 0 {

			// pfStart := i
			// fPln(obj[pfStart : pfStart+len(lookfor)])

			pvS, pvE := i+len(lookfor), 0
			pv1End, pv2End := 0, 0
			if obj[pvS] != '[' {
				pv1End = sIndex(obj[pvS:], Trait1EndV)
				pv2End = sIndex(obj[pvS:], Trait2EndV)
			} else {
				if pv1End = sIndex(obj[pvS:], Trait3EndV); pv1End >= 0 {
					pv1End++
				}
				if pv2End = sIndex(obj[pvS:], Trait4EndV); pv2End >= 0 {
					pv2End++
				}
			}

			switch {
			case pv1End != -1 && pv2End == -1:
				pvE = pv1End
			case pv1End == -1 && pv2End != -1:
				pvE = pv2End
			default:
				pvE = int(math.Min(float64(pv1End), float64(pv2End)))
			}

			valData := obj[pvS : pvS+pvE]
			// fPln(valData)

			// For Mask-JKV, only use end-leaf Mask Value
			if hashRExp.FindStringIndex(valMask) == nil {
				switch valMask {
				case `"[]"`:
					if valData[0] != '[' { // only deal with one element to one element-array
						obj = obj[:pvS] + "[" + valData + "]" + obj[pvS+pvE:] // format is needed for outcome
					}
				case `"(B)"`:
					if valData == `"true"` || valData == `"false"` {
						valData = valData[1 : len(valData)-1]
						obj = obj[:pvS] + valData + obj[pvS+pvE:]
					}
				case `"(N)"`:
					if valData[0] == '"' && valData[len(valData)-1] == '"' {
						valData = valData[1 : len(valData)-1]
					}
					if valData[0] == '.' { // deal with like ".5" format digits
						valData = "0" + valData // convert it to "0.5"
					}
					if isNumeric(valData) {
						obj = obj[:pvS] + valData + obj[pvS+pvE:]
					}
				default:
					obj = obj[:pvS] + valMask + obj[pvS+pvE:]
				}
			}

			// For Mask-JKV, only use end-leaf Mask Value
			// if hashRExp.FindStringIndex(valMask) == nil {
			// 	obj = obj[:pvS] + valMask + obj[pvS+pvE:]
			// } else {
			// 	fPln(valMask, valData)
			// }
		}
	}

	return obj
}
