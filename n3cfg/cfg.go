package n3cfg

import (
	"os"
	"os/exec"
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
