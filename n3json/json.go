package n3json

import (
	"encoding/json"
	"regexp"
	"strings"
	"sync"

	"github.com/cdutwhu/n3-util/n3err"
)

// JSONRoot :
func JSONRoot(jsonstr string) string {
	x := make(map[string]interface{})
	failOnErr("%v", json.Unmarshal([]byte(jsonstr), &x))
	for k := range x {
		return k
	}
	return ""
}

// InnerFmt :
func InnerFmt(str string) string {
	str = sTrim(str, BLANK)
	str = sReplaceAll(str, "\t", "    ")
	i := len(str) - 1
	N := 0
	if str[i] == '}' {
		for i = i - 1; i >= 0; i-- {
			if str[i] == ' ' {
				N++
				continue
			}
			break
		}
	}
	return indent(str, -N, true)
}

// MaybeArr : before using this, make sure it is valid json
func MaybeArr(str string) bool {
	return sTrimLeft(str, " \t\n\r")[0] == '['
}

// SplitArr : json doesn't need to be Formatted
func SplitArr(json string, nSpace int) []string {
	if !MaybeArr(json) {
		return nil
	}
	psGrp, peGrp := []int{}, []int{}
	lvlCnt, lvlCntPrev := 0, 0
	for i, c := range json {
		switch {
		case c == '{':
			lvlCnt++
		case c == '}':
			lvlCnt--
		}
		if lvlCnt == 1 && lvlCntPrev == 0 {
			psGrp = append(psGrp, i)
		}
		if lvlCnt == 0 && lvlCntPrev == 1 {
			peGrp = append(peGrp, i)
		}
		lvlCntPrev = lvlCnt
	}
	failOnErrWhen(len(psGrp) != len(peGrp), "%v", n3err.JSON_ARRAY_INVALID)

	// [parallel mode]
	wg := sync.WaitGroup{}
	wg.Add(len(psGrp))

	jsonGrp := make([]string, len(psGrp))
	for i, ps := range psGrp {
		pe := peGrp[i]
		// jsonGrp[i] = FmtJSON(json[ps:pe+1], nSpace) // [serial mode]

		// [parallel mode]
		go func(i, ps, pe int) {
			defer wg.Done()
			jsonGrp[i] = Fmt(json[ps:pe+1], nSpace)
		}(i, ps, pe)
	}

	// [parallel mode]
	wg.Wait()

	return jsonGrp
}

// MakeArr :
func MakeArr(jsonlist ...string) (arrstr string) {
	combine := "[\n" + sJoin(jsonlist, ",\n")
	fmtArr := indent(combine, 2, true)
	return fmtArr + "\n]"
}

// ---------------------------------------------------- //

// Merge4Async :
func Merge4Async(chGrp ...<-chan string) string {
	var jsonGrp []string
	for _, ch := range chGrp {
		jsonGrp = append(jsonGrp, <-ch)
	}
	return Merge(jsonGrp...)
}

// Merge :
func Merge(jsonGrp ...string) string {
	switch {
	case len(jsonGrp) >= 3:
		var builder strings.Builder
		for i, json := range jsonGrp {
			if i == 0 {
				p := sLastIndex(json, "}")
				builder.WriteString(sTrimRight(json[:p], " \t\r\n"))
			} else if i == len(jsonGrp)-1 {
				p := sIndex(json, "{")
				builder.WriteString(",")
				builder.WriteString(json[p+1:])
			} else {
				p1, p2 := sIndex(json, "{"), sLastIndex(json, "}")
				builder.WriteString(",")
				builder.WriteString(sTrimRight(json[p1+1:p2], " \t\r\n"))
			}
		}
		return builder.String()
	case len(jsonGrp) == 2:
		json1, json2 := jsonGrp[0], jsonGrp[1]
		p1 := sLastIndex(json1, "}")
		p2 := sIndex(json2, "{")
		return json1[:p1] + "," + json2[p2+1:]
	case len(jsonGrp) == 1:
		return jsonGrp[0]
	}
	return ""
}

// AsyncScalarSel :
func AsyncScalarSel(json, attr string) <-chan string {
	c := make(chan string)
	go func() {
		var builder strings.Builder
		builder.WriteString(fSf("{\n  \"%s\": [\n", attr))
		tag := fSf("\"%s\": ", attr)
		offset := len(tag)
		r := regexp.MustCompile(fSf(`%s.+,?\n`, tag))
		for _, l := range r.FindAllString(json, -1) {
			builder.WriteString("    ")
			l = sTrimRight(l, ",\r\n")[offset:]
			builder.WriteString(l)
			builder.WriteString(",\n")
		}
		builder.WriteString("  ]\n}")
		ret := builder.String()

		r = regexp.MustCompile(`,\n[ ]+\]`)
		pairs := r.FindAllStringIndex(ret, -1)
		failOnErrWhen(len(pairs) > 1, "%v", n3err.INTERNAL)
		if len(pairs) == 1 {
			rmPos := pairs[0][0]
			ret = ret[:rmPos] + ret[rmPos+1:]
		}
		c <- ret
	}()
	return c
}

// ScalarSelX :
func ScalarSelX(json string, attrGrp ...string) string {
	chans := make([]<-chan string, len(attrGrp))
	for i, attr := range attrGrp {
		chans[i] = AsyncScalarSel(json, attr)
	}
	return Merge4Async(chans...)
}

// ---------------------------------------------------- //

// L1Attrs : Level-1 attributes
func L1Attrs(json string) (attrs []string) {
	failP1OnErrWhen(!isJSON(json), "%v", n3err.PARAM_INVALID_JSON)
	json = Fmt(json, 2)
	r := regexp.MustCompile(`\n  "[^"]+": [\[\{"-1234567890ntf]`)
	found := r.FindAllString(json, -1)
	for _, a := range found {
		attrs = append(attrs, a[4:len(a)-4])
	}
	return
}

// Join :
func Join(jsonL, fkey, jsonR, pkey, name string) (string, bool) {
	if name == "" {
		if hasAnySuffix(pkey, "-ID", "-id", "-Id", "_ID", "_id", "_Id") {
			name = pkey[:len(pkey)-3]
		}
	}

	inputs, keys, keyTypes := []string{jsonL, jsonR}, []string{fkey, pkey}, []string{"foreign", "primary"}
	starts, ends := []int{0, 0}, []int{0, 0}
	keyLines, keyValues := []string{"", ""}, []string{"", ""}
	posGrp := [][]int{}

	for i := 0; i < 2; i++ {
		lsAttr := toGeneralSlc(L1Attrs(inputs[i]))
		failOnErrWhen(!exist(keys[i], lsAttr...), "%v: NO %s key attribute [%s]", n3err.INTERNAL, keyTypes[i], keys[i])

		r := regexp.MustCompile(fSf(`\n  "%s": .+[,]?\n`, keys[i]))
		pSEs := r.FindAllStringIndex(inputs[i], 1)
		failOnErrWhen(len(pSEs) == 0, "%v: %s key's value error", n3err.INTERNAL, keyTypes[i])
		starts[i], ends[i] = pSEs[0][0], pSEs[0][1]
		keyLines[i] = sTrim(inputs[i][starts[i]:ends[i]], ", \t\r\n")
		keyValues[i] = keyLines[i][len(fkey)+4:]

		if i == 0 {
			posGrp = pSEs
			failOnErrWhen(exist(name, lsAttr...), "%v: [%s] already exists in left json", n3err.INTERNAL, name)
		}
	}

	if keyValues[0] == keyValues[1] {
		comma := ","
		if jsonL[posGrp[0][1]] == '}' {
			comma = ""
		}
		insert := fSf(`"%s": %s%s`, name, jsonR, comma)
		str := replByPosGrp(jsonL, posGrp, []string{insert})
		return Fmt(str, 2), true
	}

	return jsonL, false
}

// ArrJoin :
func ArrJoin(jsonarrL, fkey, jsonarrR, pkey, name string) (ret string, pairs [][2]int) {
	jsonLarr := SplitArr(jsonarrL, 2)
	jsonRarr := SplitArr(jsonarrR, 2)
	joined := []string{}
	for i, jsonL := range jsonLarr {
		for j, jsonR := range jsonRarr {
			if join, ok := Join(jsonL, fkey, jsonR, pkey, name); ok {
				// fPln(ok, i, j)
				pairs = append(pairs, [2]int{i, j})
				joined = append(joined, join)
			}
		}
	}
	return MakeArr(joined...), pairs
}
