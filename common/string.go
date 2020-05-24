package common

import (
	"sort"

	eg "github.com/cdutwhu/n3-util/n3errs"
)

// HasAnyPrefix :
func HasAnyPrefix(s string, lsPrefix ...string) bool {
	FailOnErrWhen(len(lsPrefix) == 0, "%v: at least input one prefix", eg.PARAM_INVALID)
	for _, prefix := range lsPrefix {
		if sHasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// HasAnySuffix :
func HasAnySuffix(s string, lsSuffix ...string) bool {
	FailOnErrWhen(len(lsSuffix) == 0, "%v: at least input one suffix", eg.PARAM_INVALID)
	for _, suffix := range lsSuffix {
		if sHasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

// RmTailFromLast :
func RmTailFromLast(s, mark string) string {
	if i := sLastIndex(s, mark); i >= 0 {
		return s[:i]
	}
	return s
}

// RmTailFromLastN :
func RmTailFromLastN(s, mark string, iLast int) string {
	rt := s
	for i := 0; i < iLast; i++ {
		rt = RmTailFromLast(rt, mark)
	}
	return rt
}

// RmTailFromFirst :
func RmTailFromFirst(s, mark string) string {
	if i := sIndex(s, mark); i >= 0 {
		return s[:i]
	}
	return s
}

// RmTailFromFirstAny :
func RmTailFromFirstAny(s string, marks ...string) string {
	if len(marks) == 0 {
		return s
	}
	const MAX = 1000000
	var I int = MAX
	for _, mark := range marks {
		if i := sIndex(s, mark); i >= 0 && i < I {
			I = i
		}
	}
	if I != MAX {
		return s[:I]
	}
	return s
}

// RmHeadToLast :
func RmHeadToLast(s, mark string) string {
	if i := sLastIndex(s, mark); i >= 0 {
		return s[i+len(mark):]
	}
	return s
}

// RmHeadToFirst :
func RmHeadToFirst(s, mark string) string {
	segs := sSplit(s, mark)
	if len(segs) > 1 {
		return sJoin(segs[1:], mark)
	}
	return s
}

// ReplByPosGrp :
func ReplByPosGrp(s string, posGrp [][]int, newStrGrp []string) (ret string) {
	if len(posGrp) == 0 {
		return s
	}

	FailOnErrWhen(len(posGrp) != len(newStrGrp) && len(newStrGrp) != 1,
		"%v: [posGrp]-[newStrGrp] OR %v: newStrGrp can only have 1 element for filling into all posGrp",
		eg.SLICES_DIF_LEN,
		eg.SLICE_INCORRECT_COUNT)

	wrapper := make([]struct {
		posPair []int
		newStr  string
	}, len(posGrp))
	for i, pair := range posGrp {
		wrapper[i].posPair = pair
		if len(newStrGrp) == 1 {
			wrapper[i].newStr = newStrGrp[0]
		} else {
			wrapper[i].newStr = newStrGrp[i]
		}
	}
	sort.Slice(wrapper, func(i, j int) bool {
		return wrapper[i].posPair[0] < wrapper[j].posPair[0]
	})

	complement := make([][2]int, len(posGrp)+1)
	for i := 0; i < len(complement); i++ {
		if i == 0 {
			complement[i][0] = 0
			complement[i][1] = wrapper[i].posPair[0]
		} else if i == len(complement)-1 {
			complement[i][0] = wrapper[i-1].posPair[1]
			complement[i][1] = len(s)
		} else {
			complement[i][0] = wrapper[i-1].posPair[1]
			complement[i][1] = wrapper[i].posPair[0]
		}
	}

	keepStrGrp := make([]string, len(complement))
	for i := 0; i < len(keepStrGrp); i++ {
		start, end := complement[i][0], complement[i][1]
		FailOnErrWhen(end < start, "%v: [end] must be greater than [start]", eg.VAR_INVALID)
		keepStrGrp[i] = s[start:end]
	}

	for i, keep := range keepStrGrp[:len(keepStrGrp)-1] {
		ret += (keep + wrapper[i].newStr)
	}
	ret += keepStrGrp[len(keepStrGrp)-1]

	return
}

// ProjectV :
func ProjectV(strlist []string, sep, trimToL, trimFromR string) [][]string {
	nSep := 0
	for _, str := range strlist {
		if n := sCount(str, sep); n > nSep {
			nSep = n
		}
	}
	rtStrList := make([][]string, nSep+1)
	for _, str := range strlist {
		for i, s := range sSplit(str, sep) {
			if trimToL != "" {
				if fd := sIndex(s, trimToL); fd >= 0 {
					s = s[fd+1:]
				}
			}
			if trimFromR != "" {
				if fd := sLastIndex(s, trimFromR); fd >= 0 {
					s = s[:fd]
				}
			}
			rtStrList[i] = append(rtStrList[i], s)
		}
	}
	for i := 0; i < len(rtStrList); i++ {
		rtStrList[i] = ToSet(rtStrList[i]).([]string)
	}
	return rtStrList
}

// Indent :
func Indent(str string, n int, ignoreFirstLine bool) (string, bool) {
	if n == 0 {
		return str, false
	}
	S := 0
	if ignoreFirstLine {
		S = 1
	}
	lines := sSplit(str, "\n")
	if n > 0 {
		space := ""
		for i := 0; i < n; i++ {
			space += " "
		}
		for i := S; i < len(lines); i++ {
			if sTrim(lines[i], " \n\t") != "" {
				lines[i] = space + lines[i]
			}
		}
	} else {
		for i := S; i < len(lines); i++ {
			if len(lines[i]) == 0 { //                                         ignore empty string line
				continue
			}
			if len(lines[i]) <= -n || sTrimLeft(lines[i][:-n], " ") != "" { // cannot be indented as <n>, give up indent
				return str, false
			}
			lines[i] = lines[i][-n:]
		}
	}
	return sJoin(lines, "\n"), true
}
