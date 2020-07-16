package cfg

import (
	"os"
	"os/exec"
	"regexp"
)

// GitVer :
func GitVer() string {
	tag := GitTag()
	r := regexp.MustCompile(`^v[0-9]+\.[0-9]+\.[0-9]+$`)
	if r.MatchString(tag) {
		return tag
	}
	return ""
}

// GitTag :
func GitTag() string {
	// check git existing
	_, oriWD := prepare("git")
	os.Chdir(oriWD) // under .git project dir to get `git tag`

	// run git
	cmdstr := "git tag"
	cmd := exec.Command("bash", "-c", cmdstr)
	output, err := cmd.Output()
	failOnErr("cmd.Output() error @ %v", err)
	outstr := sTrim(string(output), " \n\t")
	if outstr == "" {
		return ""
	}
	lines := sSplit(outstr, "\n")
	return lines[len(lines)-1]
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
					field := cfgElem.Field(i)
					if oriVal, ok := field.Interface().(string); ok && sContains(oriVal, key) {
						if repVal, ok := value.(string); ok {
							field.SetString(sReplaceAll(oriVal, key, repVal))
						} else {
							field.Set(vof(value))
						}
					}
					// go into struct element
					if field.Kind() == typSTRUCT {
						for j, nFieldSub := 0, field.NumField(); j < nFieldSub; j++ {
							fieldSub := field.Field(j)
							if oriVal, ok := fieldSub.Interface().(string); ok && sContains(oriVal, key) {
								if repVal, ok := value.(string); ok {
									fieldSub.SetString(sReplaceAll(oriVal, key, repVal))
								} else {
									fieldSub.Set(vof(value))
								}
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
