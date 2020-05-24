package external

import (
	"os"

	eg "github.com/cdutwhu/n3-util/n3errs"
)

// Prepare : // append `defer func() { os.Chdir(oriWD) }()` after this function
func Prepare(exe string, wkDirs ...string) (cWD, oriWD string, err error) {
	oriWD, err = os.Getwd()
	failOnErr("%v", err)
	wkDirs = append(wkDirs, "./", "../", "../../")
	for _, cWD = range wkDirs {
		cWD = sTrimRight(cWD, "/") + "/"
		if _, err = os.Stat(cWD + exe); err == nil {
			failOnErr("%v", os.Chdir(cWD))
			cWD, err = os.Getwd()
			failOnErr("%v", err)
			return cWD, oriWD, nil
		}
	}
	failOnErr("%v: %s", eg.FILE_NOT_FOUND, exe)
	return "", "", nil
}
