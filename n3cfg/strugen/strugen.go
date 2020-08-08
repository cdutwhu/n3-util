package strugen

import (
	"io/ioutil"
	"os/user"
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
func GenStruct(tomlFile, struName, pkgName, struFile string) {

	failP1OnErrWhen(!sHasSuffix(tomlFile, ".toml"), "%v @tomlFile", n3err.PARAM_INVALID)
	tomlFile, err := filepath.Abs(tomlFile)
	failP1OnErr("%v", err)
	bytes, err := ioutil.ReadFile(tomlFile)
	failP1OnErr("%v", err)

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
	return
}

// AddCfg2Bank : echo 'password' | sudo -S env "PATH=$PATH" go test -v ./ -run TestAddCfg2Bank
func AddCfg2Bank(funcOSUser, tomlFile, cfgName, pkgName string) string {
	enableLog2F(true, logfile)
	cfgName, pkgName = sTitle(cfgName), sToLower(pkgName)
	dir, _ := callerSrc()
	file := filepath.Dir(dir) + fSf("/%s/%s/%s.go", "bank", pkgName, cfgName) // cfg struct Name as to be go fileName

	if funcOSUser == "" {
		user, err := user.Current()
		failOnErr("%v", err)
		funcOSUser = user.Name
		file = sReplace(file, "/root/", "/home/"+funcOSUser+"/", 1)
	}

	logger("ready to generate: %v", file)
	GenStruct(tomlFile, cfgName, pkgName, file)
	logger("finish generating: %v", file)

	// file LIKE `/home/qmiao/go/pkg/mod/github.com/cdutwhu/n3-util@v0.2.27/n3cfg/bank/s2jsvr/Config.go`
	pkgmark := "/go/pkg/mod/"
	if sContains(file, pkgmark) {
		fullpkg := filepath.Dir(sSplit(file, pkgmark)[1])
		logger("generated package path: %v", fullpkg)
		pos := rxMustCompile(`@[^/]+/`).FindAllStringIndex(fullpkg, -1)
		pkg := replByPosGrp(fullpkg, pos, []string{""}, 0, 1)
		logger("generated package: %v", pkg)
		return pkg
	}
	return file
}
