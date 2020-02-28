package common

import "sort"

// HasAnyPrefix :
func HasAnyPrefix(s string, lsPrefix ...string) bool {
	FailOnErrWhen(len(lsPrefix) == 0, "%v", fEf("at least one prefix for input"))
	for _, prefix := range lsPrefix {
		if sHasPrefix(s, prefix) {
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
		return s[i+len(mark) : len(s)]
	}
	return s
}

// RmHeadToFirst :
func RmHeadToFirst(s, mark string) string {
	segs := sSpl(s, mark)
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
		"%v",
		fEf("posGrp's length must be equal to newStrGrp's length OR newStrGrp only has 1 element for filling into all posGrp"))

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
		FailOnErrWhen(end < start, "%v", fEf("end must be greater than start"))
		keepStrGrp[i] = s[start:end]
	}

	for i, keep := range keepStrGrp[:len(keepStrGrp)-1] {
		ret += (keep + wrapper[i].newStr)
	}
	ret += keepStrGrp[len(keepStrGrp)-1]

	return
}
