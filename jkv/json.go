package jkv

import (
	"regexp"
	"strings"
	"sync"

	cmn "github.com/cdutwhu/json-util/common"
)

// MaybeJSONArr : before using this, make sure it is valid json
func MaybeJSONArr(str string) bool {
	return sTrimLeft(str, " \t\n\r")[0] == '['
}

// SplitJSONArr : json doesn't need to be Formatted
func SplitJSONArr(json string, nSpace int) []string {
	if !MaybeJSONArr(json) {
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
	cmn.FailOnErrWhen(len(psGrp) != len(peGrp), "%v", fEf("Fatal, is valid JSON array?"))

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
			jsonGrp[i] = FmtJSON(json[ps:pe+1], nSpace)
		}(i, ps, pe)
	}

	// [parallel mode]
	wg.Wait()

	return jsonGrp
}

// MakeJSONArr :
func MakeJSONArr(jsonlist ...string) (arrstr string) {
	combine := "[\n" + sJoin(jsonlist, ",\n")
	fmtArr, _ := Indent(combine, 2, true)
	return fmtArr + "\n]"
}

// ****************************************************** //

// JSONMerge4Async :
func JSONMerge4Async(chGrp ...<-chan string) string {
	var jsonGrp []string
	for _, ch := range chGrp {
		jsonGrp = append(jsonGrp, <-ch)
	}
	return JSONMerge(jsonGrp...)
}

// JSONMerge :
func JSONMerge(jsonGrp ...string) string {
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

// AsyncJSONScalarSel :
func AsyncJSONScalarSel(json, attr string) <-chan string {
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
		cmn.FailOnErrWhen(len(pairs) > 1, "%v", fEf("Error"))
		if len(pairs) == 1 {
			rmPos := pairs[0][0]
			ret = ret[:rmPos] + ret[rmPos+1:]
		}
		c <- ret
	}()
	return c
}

// JSONScalarSelX :
func JSONScalarSelX(json string, attrGrp ...string) string {
	chans := make([]<-chan string, len(attrGrp))
	for i, attr := range attrGrp {
		chans[i] = AsyncJSONScalarSel(json, attr)
	}
	return JSONMerge4Async(chans...)
}
