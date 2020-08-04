package n3cfg

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
)

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
	lines := sSplit(outstr, "\n")
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
	lines := sSplit(string(bytes), "\n")
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

// ----------------------------------- //

// New :
func New(cfg interface{}, mReplExpr map[string]string, cfgPaths ...string) string {
	defer func() { mux.Unlock() }()
	mux.Lock()
	for _, f := range cfgPaths {
		if _, e := os.Stat(f); e == nil {
			initCfg(f, cfg, mReplExpr)
			return f
		}
	}
	return ""
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
		"[IP]":   localIP(),
		"[GV]":   ver,
		"[PATH]": abs,
	})

	mRepl := make(map[string]interface{})
	for k, v := range mReplExpr {
		mRepl[k] = EvalCfgValue(cfg, v)
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

// InitEnvVar : initialize the global variables
func InitEnvVar(cfg interface{}, mReplExpr map[string]string, key string, cfgPaths ...string) bool {
	if New(cfg, mReplExpr, append(cfgPaths, "./config.toml", "./config/config.toml")...) == "" {
		return false
	}
	struct2Env(key, cfg)
	return true
}
