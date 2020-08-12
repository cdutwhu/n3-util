package strugen

import (
	"io/ioutil"
	"path/filepath"

	"github.com/cdutwhu/n3-util/n3err"
)

func scanToml(tomllines []string) (slAttrs, grpAttrs []string) {
	chkStartGroup := func(line string) bool {
		ln := sTrimLeft(line, " \t")
		return sHasPrefix(ln, "[")
	}
	attrSingle := true
	for _, line := range tomllines {
		if attrSingle {
			if sContains(line, "=") {
				attr := sTrim(rmTailFromFirst(line, "="), " \t")
				slAttrs = append(slAttrs, attr)
			}
			if chkStartGroup(line) {
				attrSingle = false
			}
		}
		if chkStartGroup(line) {
			ln := rmTailFromFirst(line, "]")
			attr := rmHeadToFirst(ln, "[")
			grpAttrs = append(grpAttrs, attr)
		}
	}
	return
}

func attrsRange(tomllines []string) map[string][2]int {
	mGrpPos := make(map[string][2]int)
	slAttrs, grpAttrs := scanToml(tomllines)
	if len(slAttrs) > 0 {
		mGrpPos[""] = [2]int{0, len(slAttrs) - 1}
	}
	chkEndGroup := func(line string) bool {
		ln := sTrim(line, " \t")
		return hasAnyPrefix(ln, "[", "#") || ln == ""
	}
	for _, gattr := range grpAttrs {
		found := false
		for i, line := range tomllines {
			if sHasPrefix(sTrim(line, " \t"), "["+gattr+"]") {
				found = true
				mGrpPos[gattr] = [2]int{i + 1, -1}
				continue
			}
			if found {
				if chkEndGroup(line) {
					mGrpPos[gattr] = [2]int{mGrpPos[gattr][0], i - 1}
					break
				}
			}
		}
	}
	return mGrpPos
}

// value contains pure '[*]', the type is 'interface{}'
func attrTypes(tomllines []string, grpAttr string) map[string]string {
	mAttrType := make(map[string]string)
	mGrpPos := attrsRange(tomllines)
	rng := mGrpPos[grpAttr]
	start, end := rng[0], rng[1]
	for i := start; i <= end; i++ {
		ln := sTrim(tomllines[i], " \t")
		attr := sTrim(rmTailFromFirst(ln, "="), " \t")
		value := sTrim(rmHeadToFirst(ln, "="), " \t")
		switch {
		case sCount(value, "\"[") == 1 && sCount(value, "]\"") == 1:
			mAttrType[attr] = "interface{}"
		case isNumeric(value) && !sContains(value, "."):
			mAttrType[attr] = "int"
		case isNumeric(value) && sContains(value, "."):
			mAttrType[attr] = "float64"
		case value == "true" || value == "false":
			mAttrType[attr] = "bool"
		default:
			mAttrType[attr] = "string"
		}
	}
	return mAttrType
}

// GenStruct :
func GenStruct(tomlFile, struName, pkgName, struFile string) bool {
	warnP1OnErrWhen(!sHasSuffix(tomlFile, ".toml"), "%v @tomlFile", n3err.PARAM_INVALID)

	tomlFile, err := filepath.Abs(tomlFile)
	if err != nil {
		warnP1OnErr("%v", err)
		return false
	}

	bytes, err := ioutil.ReadFile(tomlFile)
	if err != nil {
		warnP1OnErr("%v", err)
		return false
	}

	fname := sReplaceAll(filepath.Base(tomlFile), ".toml", "")
	dir := filepath.Dir(tomlFile)
	if struName == "" {
		struName = sToUpper(string(fname[0])) + fname[1:]
	}
	if pkgName == "" {
		pkgName = filepath.Base(dir)
	}
	if struFile == "" {
		struFile = dir + "/" + fname + ".go"
	}

	failP1OnErrWhen(!ucIsUpper(rune(struName[0])), "%v @struName", n3err.PARAM_INVALID)

	lines := splitLn(string(bytes))
	struStr := ""
	if pkgName != "" {
		struStr += fSf("package %s\n\n", pkgName)
	}

	struStr += fSf("// %s : AUTO Created From %s\n", struName, tomlFile)
	struStr += fSf("type %s struct {\n", struName)
	for k, v := range attrTypes(lines, "") { // root type is ""
		struStr += fSf("    %s %s\n", k, v) // 4 space
	}
	_, attrs2 := scanToml(lines)
	for _, attr := range attrs2 {
		struStr += fSf("    %s struct {\n", attr)
		for k, v := range attrTypes(lines, attr) {
			struStr += fSf("        %s %s\n", k, v) // 8 space
		}
		struStr += fSln("    }") // 4 space
	}
	struStr += fSln("}")

	mustWriteFile(struFile, []byte(struStr))
	return true
}
