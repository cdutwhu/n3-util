package n3cfg

import "io/ioutil"

// RmFileAttrL1 :
func RmFileAttrL1(infile, outfile string, attrs ...string) string {
	bytes, err := ioutil.ReadFile(infile)
	failP1OnErr("%v", err)
	// if sHasSuffix(outfile, ".toml") {
	// 	outfile = outfile[:len(outfile)-5]
	// }
	// outfile = rmTailFromLast(infile, "/") + "/" + outfile + ".toml"
	if !sHasSuffix(outfile, ".toml") {
		outfile += ".toml"
	}
	mustWriteFile(outfile, []byte(rmAttrL1(string(bytes), attrs...)))
	return outfile
}

// rmAttrL1 :
func rmAttrL1(toml string, attrs ...string) string {
	chkStart := func(line, attr string) bool {
		return sHasPrefix(line, "["+attr+"]")
	}
	chkEnd := func(line string) bool {
		ln := sTrim(line, " \t")
		return ln == "" || sHasPrefix(ln, "#") || sHasPrefix(ln, "[")
	}

	chkEndOfSingle := func(line string) bool {
		ln := sTrimLeft(line, " \t")
		return sHasPrefix(ln, "[")
	}

	pairs, rmflag, attrSingle := [][2]int{}, false, true
	lines := sSplit(toml, "\n")
NEXT1:
	for i, line := range lines {
		for _, attr := range attrs {

			// ------------------------- //
			if attrSingle {
				if ln := sTrim(rmTailFromFirst(line, "="), " \t"); ln == attr {
					pairs = append(pairs, [2]int{i, i})
				}
				if chkEndOfSingle(line) {
					attrSingle = false
				}
			}
			// ------------------------- //

			if chkStart(line, attr) {
				pairs = append(pairs, [2]int{i, -1})
				rmflag = true
				continue NEXT1
			}
			if rmflag && chkEnd(line) {
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
