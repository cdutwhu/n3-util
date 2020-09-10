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
	if rng, ok := mGrpPos[grpAttr]; ok {
		for i := rng[0]; i <= rng[1]; i++ {
			ln := sTrim(tomllines[i], " \t")
			attr := sTrim(rmTailFromFirst(ln, "="), " \t")
			val := sTrim(rmHeadToFirst(ln, "="), " \t")
			switch {
			case sCount(val, "\"[") == 1 && sCount(val, "]\"") == 1:
				mAttrType[attr] = "interface{}"
			case isNumeric(val) && !sContains(val, "."):
				mAttrType[attr] = "int"
			case isNumeric(val) && sContains(val, "."):
				mAttrType[attr] = "float64"
			case val == "true" || val == "false":
				mAttrType[attr] = "bool"
			case sHasPrefix(val, "[") && sHasSuffix(val, "]"):
				first := sSplit(val[1:len(val)-1], ",")[0]
				switch {
				case isNumeric(first) && !sContains(first, "."):
					mAttrType[attr] = "[]int"
				case isNumeric(first) && sContains(first, "."):
					mAttrType[attr] = "[]float64"
				case first == "true" || first == "false":
					mAttrType[attr] = "[]bool"
				default:
					mAttrType[attr] = "[]string"
				}
			default:
				mAttrType[attr] = "string"
			}
		}
	}
	return mAttrType
}

// GenStruct :
func GenStruct(tomlFile, struName, pkgName, struFile string) bool {
	warnP1OnErrWhen(!sHasSuffix(tomlFile, ".toml"), "%v @tomlFile", n3err.PARAM_INVALID)

	tomlFile, err := filepath.Abs(tomlFile)
	if warnP1OnErr("%v", err) != nil {
		return false
	}

	bytes, err := ioutil.ReadFile(tomlFile)
	if warnP1OnErr("%v", err) != nil {
		return false
	}

	fname := sReplaceAll(filepath.Base(tomlFile), ".toml", "")
	dir := filepath.Dir(tomlFile)
	if struName == "" {
		struName = sToUpper(string(fname[0])) + fname[1:]
	}
	failP1OnErrWhen(!ucIsUpper(rune(struName[0])), "%v @struName", n3err.PARAM_INVALID)

	addPkg := true
	if pkgName == "" {
		pkgName = filepath.Base(dir)
	}
	if struFile == "" {
		struFile = dir + "/" + fname + ".go"
	}
	if bytes, err := ioutil.ReadFile(struFile); err == nil && sHasPrefix(string(bytes), "package ") {
		addPkg = false
	}

	lines := splitLn(string(bytes) + "\n")
	struStr := ""
	if addPkg {
		struStr += fSf("package %s\n\n", pkgName)
		struStr += `import "github.com/cdutwhu/n3-util/n3cfg"` + "\n\n"
	}

	struStr += fSf("// %s : AUTO Created From %s\n", struName, tomlFile)
	struStr += fSf("type %s struct {\n", struName)
	for k, v := range attrTypes(lines, "") { // root type is ""
		struStr += fSf("\t%s %s\n", k, v)
	}
	_, attrs2 := scanToml(lines)
	for _, attr := range attrs2 {
		struStr += fSf("\t%s struct {\n", attr)
		for k, v := range attrTypes(lines, attr) {
			struStr += fSf("\t\t%s %s\n", k, v)
		}
		struStr += fSln("\t}")
	}
	struStr += fSln("}")

	mustAppendFile(struFile, []byte(struStr), true)
	return true
}

// GenNewCfg :
func GenNewCfg(struFile string) bool {
	bytes, err := ioutil.ReadFile(struFile)
	if err != nil {
		return false
	}
	cont := string(bytes)
	struNames := []string{}
	r := rxMustCompile(`type [A-Za-z0-9_]+ struct {`)
	offset1, offset2 := len("type "), len(" struct {")
	for _, find := range r.FindAllString(cont, -1) {
		struNames = append(struNames, find[offset1:len(find)-offset2])
	}
	fPln(struNames)
	if sContains(cont, "func NewCfg(") {
		return false
	}

	src := `// NewCfg :` + "\n"
	src += `func NewCfg(cfgStruName string, mReplExpr map[string]string, cfgPaths ...string) interface{} {` + "\n"
	src += `	var cfg interface{}` + "\n"
	src += `	switch cfgStruName {` + "\n"
	for _, cfgname := range struNames {
		src += fSf("\tcase \"%[1]s\":\n\t\tcfg = &%[1]s{}\n", cfgname)
	}
	src += "\tdefault:\n\t\treturn nil\n"
	src += "\t}\n"
	src += `	return n3cfg.InitEnvVar(cfg, mReplExpr, cfgStruName, cfgPaths...)` + "\n"
	src += `}` + "\n"
	
	mustAppendFile(struFile, []byte(src), true)
	return false
}
