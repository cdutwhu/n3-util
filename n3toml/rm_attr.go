package n3toml

import "io/ioutil"

// RmFileAttrL1 :
func RmFileAttrL1(inputfile, outname string, attrs ...string) string {
	bytes, err := ioutil.ReadFile(inputfile)
	failOnErr("%v", err)
	outname = sTrimRight(outname, ".toml")
	outfile := rmTailFromLast(inputfile, "/") + "/" + outname + ".toml"
	mustWriteFile(outfile, []byte(RmAttrL1(string(bytes), attrs...)))
	return outfile
}

// RmAttrL1 :
func RmAttrL1(toml string, attrs ...string) string {
	chkStart := func(line, attr string) bool {
		return sHasPrefix(line, "["+attr+"]")
	}
	chkEnd := func(line, attr string) bool {
		ln := sTrim(line, " \t")
		return ln == "" || sHasPrefix(ln, "#") || sHasPrefix(ln, "[")
	}

	pairs, rmflag := [][2]int{}, false
	lines := sSplit(toml, "\n")
NEXT1:
	for i, line := range lines {
		for _, attr := range attrs {
			if chkStart(line, attr) {
				pairs = append(pairs, [2]int{i, -1})
				rmflag = true
				continue NEXT1
			}
			if rmflag && chkEnd(line, attr) {
				pairs[len(pairs)-1][1] = i - 1
				rmflag = false
			}
		}
	}

	remain := []string{}
NEXT2:
	for i, line := range lines {
		for _, pair := range pairs {
			start, end := pair[0], pair[1]
			if i >= start && i <= end {
				continue NEXT2
			}
		}
		remain = append(remain, line)
	}

	ret := sJoin(remain, "\n")
AGAIN:
	if sContains(ret, "\n\n\n") {
		ret = sReplaceAll(ret, "\n\n\n", "\n\n")
		goto AGAIN
	}
	return ret
}
