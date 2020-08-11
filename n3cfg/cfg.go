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
