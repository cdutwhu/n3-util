package attrim

import "io/ioutil"

// ------------------------- Remove ------------------------- //

// attrL1Rm :
func attrL1Rm(toml string, attrs ...string) string {
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

	pairs, flagRm, flagSA := [][2]int{}, false, true
	lines := splitLn(toml)
NEXT1:
	for i, line := range lines {
		for _, attr := range attrs {

			// ------------------------- //
			if flagSA {
				if ln := sTrim(rmTailFromFirst(line, "="), " \t"); ln == attr {
					pairs = append(pairs, [2]int{i, i})
				}
				if chkEndOfSingle(line) {
					flagSA = false
				}
			}
			// ------------------------- //

			if chkStart(line, attr) {
				pairs = append(pairs, [2]int{i, -1})
				flagRm = true
				continue NEXT1
			}
			if flagRm && chkEnd(line) {
				pairs[len(pairs)-1][1] = i - 1
				flagRm = false
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
	if !sHasSuffix(ret, "\n") {
		ret += "\n"
	}
	return ret
}

// RmCfgAttrL1 :
func RmCfgAttrL1(infile, outfile string, attrs ...string) string {
	bytes, err := ioutil.ReadFile(infile)
	failP1OnErr("%v", err)
	if !sHasSuffix(outfile, ".toml") {
		outfile += ".toml"
	}
	mustWriteFile(outfile, []byte(attrL1Rm(string(bytes), attrs...)))
	return outfile
}

// ------------------------- Select ------------------------- //

// attrL1Sel :
func attrL1Sel(toml string, attrs ...string) string {
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

	pairs, flagSel, flagSA := [][2]int{}, false, true
	lines := splitLn(toml)
NEXT1:
	for i, line := range lines {
		for _, attr := range attrs {

			// ------------------------- //
			if flagSA {
				if ln := sTrim(rmTailFromFirst(line, "="), " \t"); ln == attr {
					pairs = append(pairs, [2]int{i, i})
				}
				if chkEndOfSingle(line) {
					flagSA = false
				}
			}
			// ------------------------- //

			if chkStart(line, attr) {
				pairs = append(pairs, [2]int{i, -1})
				flagSel = true
				continue NEXT1
			}
			if flagSel && chkEnd(line) {
				pairs[len(pairs)-1][1] = i - 1
				flagSel = false
			}
		}
	}

	selection := []string{}
	for i, line := range lines {
		for _, pair := range pairs {
			start, end := pair[0], pair[1]
			if i >= start && i <= end {
				selection = append(selection, line)
			}
		}
	}

	ret := sJoin(selection, "\n")
	ret = sReplaceAll(ret, "\n[", "\n\n[") // add blank lines above each group attribute
AGAIN:
	if sContains(ret, "\n\n\n") {
		ret = sReplaceAll(ret, "\n\n\n", "\n\n")
		goto AGAIN
	}
	if !sHasSuffix(ret, "\n") {
		ret += "\n"
	}
	return ret
}

// SelCfgAttrL1 :
func SelCfgAttrL1(infile, outfile string, attrs ...string) string {
	bytes, err := ioutil.ReadFile(infile)
	failP1OnErr("%v", err)
	if !sHasSuffix(outfile, ".toml") {
		outfile += ".toml"
	}
	mustWriteFile(outfile, []byte(attrL1Sel(string(bytes), attrs...)))
	return outfile
}
