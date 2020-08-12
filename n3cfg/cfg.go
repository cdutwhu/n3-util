package n3cfg

import (
	"encoding/json"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/cdutwhu/n3-util/n3cfg/strugen"
)

// PrjName :
func PrjName() string {
	const check = "/.git"
NEXT:
	for i := 1; i < 64; i++ {
		for _, ln := range splitLn(trackCaller(i)) {
			if sHasPrefix(ln, "/") {
				ln = rmTailFromLast(ln, ":")
			AGAIN:
				dir := filepath.Dir(ln)
				if dir == "/" {
					continue NEXT
				}
				_, err := os.Stat(dir + check)
				if os.IsNotExist(err) {
					ln = dir
					goto AGAIN
				} else {
					return filepath.Base(dir)
				}
			}
		}
	}
	return ""
}

// GitVer :
func GitVer() (ver string, err error) {
	tag, err := GitTag()
	if err != nil {
		return "", err
	}
	if r := rMustCompile(`^v[0-9]+\.[0-9]+\.[0-9]+$`); r.MatchString(tag) {
		return tag, nil
	}
	return "", nil
}

// GitTag :
func GitTag() (tag string, err error) {
	defer func() {
		if r := recover(); r != nil {
			tag, err = "", fEf("%v", r)
		}
	}()

	// check git existing
	_, oriWD := prepare("git") // maybe invoke panic
	os.Chdir(oriWD)            // under .git project dir to get `git tag`

	// run git
	cmd := exec.Command("bash", "-c", "git tag")
	output, err := cmd.Output()
	failOnErr("cmd.Output() error @ %v", err)
	outstr := sTrim(string(output), " \n\t")
	if outstr == "" {
		return "", nil
	}
	lines := splitLn(outstr)
	if len(lines) >= 1 {
		return lines[len(lines)-1], nil
	}
	return lines[0], nil
}

// Modify : only 2 levels struct variable could be modified. that is enough for config
func Modify(cfg interface{}, mRepl map[string]interface{}) interface{} {
	if mRepl == nil || len(mRepl) == 0 {
		return cfg
	}
	if vof(cfg).Kind() == typPTR {
		if cfgElem := vof(cfg).Elem(); cfgElem.Kind() == typSTRUCT {
			for i, nField := 0, cfgElem.NumField(); i < nField; i++ {
				for key, value := range mRepl {
					var ivalue interface{} = value
					repVal, isstrValue := value.(string)
					field := cfgElem.Field(i)
					if oriVal, ok := field.Interface().(string); ok && sContains(oriVal, key) {
						if isstrValue {
							ivalue = sReplaceAll(oriVal, key, repVal)
						}
						field.Set(vof(ivalue))
					}
					// go into struct element
					if field.Kind() == typSTRUCT {
						for j, nFieldSub := 0, field.NumField(); j < nFieldSub; j++ {
							fieldSub := field.Field(j)
							if oriVal, ok := fieldSub.Interface().(string); ok && sContains(oriVal, key) {
								if isstrValue {
									ivalue = sReplaceAll(oriVal, key, repVal)
								}
								fieldSub.Set(vof(ivalue))
							}
						}
					}
				}
			}
			return cfg
		}
	}
	failP1OnErr("%v", fEf("input cfg MUST be struct pointer"))
	return nil
}

// EvalCfgValue :
func EvalCfgValue(cfg interface{}, key string) interface{} {
	bytes, err := json.MarshalIndent(cfg, "", "\t")
	failP1OnErr("%v", err)
	lines := splitLn(string(bytes))
	if sCount(key, ".") == 0 {
		for _, ln := range lines {
			if sHasPrefix(ln, fSf("\t\"%s\":", key)) {
				sval := sTrim(sSplit(ln, ": ")[1], ",\"")
				switch {
				case isNumeric(sval) && !sContains(sval, "."):
					ret, _ := strconv.ParseInt(sval, 10, 64)
					return int(ret)
				case isNumeric(sval) && sContains(sval, "."):
					ret, _ := strconv.ParseFloat(sval, 64)
					return ret
				case sval == "true" || sval == "false":
					ret, _ := strconv.ParseBool(sval)
					return ret
				default:
					return sval
				}
			}
		}
	} else if sCount(key, ".") == 1 {
		ss := sSplit(key, ".")
		part1, part2 := ss[0], ss[1]
	NEXT:
		for i, ln1 := range lines {
			if sHasPrefix(ln1, fSf("\t\t\"%s\":", part2)) {
				for j := i - 1; j >= 0; j-- {
					ln2 := lines[j]
					if sHasPrefix(ln2, "\t\"") {
						if sHasPrefix(ln2, fSf("\t\"%s\":", part1)) {
							sval := sTrim(sSplit(ln1, ": ")[1], ",\"")
							switch {
							case isNumeric(sval) && !sContains(sval, "."):
								ret, _ := strconv.ParseInt(sval, 10, 64)
								return int(ret)
							case isNumeric(sval) && sContains(sval, "."):
								ret, _ := strconv.ParseFloat(sval, 64)
								return ret
							case sval == "true" || sval == "false":
								ret, _ := strconv.ParseBool(sval)
								return ret
							default:
								return sval
							}
						}
						continue NEXT
					}
				}
			}
		}
	}
	return nil
}

// ------------------------------------------------------------------------------- //

// InitEnvVar : initialize the global variables
func InitEnvVar(cfg interface{}, mReplExpr map[string]string, key string, cfgPaths ...string) interface{} {
	cfg = New(cfg, mReplExpr, append(cfgPaths, "./config.toml")...)
	if cfg == nil {
		return nil
	}
	struct2Env(key, cfg)
	return cfg
}

// New :
func New(cfg interface{}, mReplExpr map[string]string, cfgPaths ...string) interface{} {
	defer func() { mux.Unlock() }()
	mux.Lock()
	for _, f := range cfgPaths {
		if _, e := os.Stat(f); e == nil {
			return initCfg(f, cfg, mReplExpr)
		}
	}
	return nil
}

func initCfg(fpath string, cfg interface{}, mReplExpr map[string]string) interface{} {
	_, e := toml.DecodeFile(fpath, cfg)
	failPnOnErr(2, "%v", e)
	abs, e := filepath.Abs(fpath)
	failOnErr("%v", e)
	home, e := os.UserHomeDir()
	failOnErr("%v", e)
	ver, e := GitVer()
	failOnErr("%v", e)
	cfg = Modify(cfg, map[string]interface{}{
		"~":      home,
		"[DATE]": time.Now().Format("2006-01-02"),
		"[PATH]": abs,
		"[IP]":   localIP(),
		"[PRJ]":  PrjName(),
		"[VER]":  ver,
	})

	mRepl := make(map[string]interface{})
	for k, v := range mReplExpr {
		value := EvalCfgValue(cfg, v)
		if value != nil {
			mRepl[k] = value
		} else {
			mRepl[k] = v
		}
	}
	return Modify(cfg, mRepl)
}

// Save :
func Save(fpath string, cfg interface{}) {
	if !sHasSuffix(fpath, ".toml") {
		fpath += ".toml"
	}
	f, e := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	failP1OnErr("%v", e)
	defer f.Close()
	failP1OnErr("%v", toml.NewEncoder(f).Encode(cfg))
}

// --------------------------------------------------------------------------------------------------- //

// Register : echo 'password' | sudo -S env "PATH=$PATH" go test -v -count=1 ./ -run TestRegister
func Register(funcOSUser, tomlFile, prjName, pkgName string) (bool, string) {
	enableLog2F(true, logfile)
	if funcOSUser == "" {
		user, err := user.Current()
		failOnErr("%v", err)
		funcOSUser = user.Name
	}

	pkgName = sToLower(pkgName)
	dir, _ := callerSrc()
	n3cfgDir := dir                                                     // filepath.Dir(dir)
	n3cfgDir = sReplace(n3cfgDir, "/root/", "/home/"+funcOSUser+"/", 1) // sudo root go pkg --> input OS-user go pkg
	file := n3cfgDir + fSf("/cache/%s/%s/Config.go", prjName, pkgName)  // cfg struct Name as to be go fileName

	logger("ready to generate: %v", file)
	if !strugen.GenStruct(tomlFile, "Config", pkgName, file) {
		return false, ""
	}
	logger("finish generating: %v", file)

	// file LIKE `/home/qmiao/go/pkg/mod/github.com/cdutwhu/n3-util@v0.2.27/n3cfg/bank/s2jsvr/Config.go`
	pkgmark := "/go/pkg/mod/"
	if sContains(file, pkgmark) {
		fullpkg := filepath.Dir(sSplit(file, pkgmark)[1])
		logger("generated package path: %v", fullpkg)
		pos := rxMustCompile(`@[^/]+/`).FindAllStringIndex(fullpkg, -1)
		pkg := replByPosGrp(fullpkg, pos, []string{""}, 0, 1)
		logger("generated package: %v", pkg)
		// make necessary functions for using
		mkFuncs(pkg, prjName, pkgName, n3cfgDir)
		return true, pkg
	}
	return false, file
}

func mkFuncs(impt, prj, pkg, fnDir string) {
	pkg = sToLower(pkg)
	CfgFnFile := fnDir + "/auto_" + prj + "_" + pkg + ".go"

	prj = sReplaceAll(prj, "-", "")
	prj = sReplaceAll(prj, " ", "")
	pkg = sReplaceAll(pkg, "-", "")
	pkg = sReplaceAll(pkg, " ", "")

	prj, pkg = sTitle(prj), sTitle(pkg)
	fnNewCfg := `New` + prj + pkg
	fnToEnvVar := `ToEnvVar` + prj + pkg
	fnFromEnvVar := `FromEnvVar` + prj + pkg

	NewCfgSrc := `package n3cfg` + "\n\n"
	NewCfgSrc += `import auto "` + impt + `"` + "\n"
	NewCfgSrc += `import "os"` + "\n\n"
	NewCfgSrc += `func ` + fnNewCfg + `(mReplExpr map[string]string, cfgPaths ...string) *auto.Config {` + "\n"
	NewCfgSrc += `    defer func() { mux.Unlock() }()` + "\n"
	NewCfgSrc += `    mux.Lock()` + "\n"
	NewCfgSrc += `    cfg := &auto.Config{}` + "\n"
	NewCfgSrc += `    for _, f := range cfgPaths {` + "\n"
	NewCfgSrc += `        if _, e := os.Stat(f); e == nil {` + "\n"
	NewCfgSrc += `            return initCfg(f, cfg, mReplExpr).(*auto.Config)` + "\n"
	NewCfgSrc += `        }` + "\n"
	NewCfgSrc += `    }` + "\n"
	NewCfgSrc += `    return nil` + "\n"
	NewCfgSrc += `}` + "\n\n"
	NewCfgSrc += `// -------------------------------- //` + "\n\n"
	NewCfgSrc += `func ` + fnToEnvVar + `(mReplExpr map[string]string, key string, cfgPaths ...string) *auto.Config {` + "\n"
	NewCfgSrc += `    cfg := ` + fnNewCfg + `(mReplExpr, append(cfgPaths, "./config.toml")...)` + "\n"
	NewCfgSrc += `    if cfg == nil {` + "\n"
	NewCfgSrc += `        return nil` + "\n"
	NewCfgSrc += `    }` + "\n"
	NewCfgSrc += `    struct2Env(key, cfg)` + "\n"
	NewCfgSrc += `    return cfg` + "\n"
	NewCfgSrc += `}` + "\n\n"
	NewCfgSrc += `// -------------------------------- //` + "\n\n"
	NewCfgSrc += `func ` + fnFromEnvVar + `(key string) *auto.Config {` + "\n"
	NewCfgSrc += `    return env2Struct(key, &auto.Config{}).(*auto.Config)` + "\n"
	NewCfgSrc += `}` + "\n\n"

	mustWriteFile(CfgFnFile, []byte(NewCfgSrc))
}
