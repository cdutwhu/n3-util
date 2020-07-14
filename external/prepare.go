package external

import (
	"os"

	eg "github.com/cdutwhu/n3-util/n3errs"
)

// prepare : // append `defer func() { os.Chdir(oriWD) }()` after this function
func prepare(exe string, wkDirs ...string) (cWD, oriWD string) {
	oriWD, err := os.Getwd()
	failOnErr("%v", err)
	for _, cWD = range append(wkDirs, "./", "../", "../../") {
		cWD = sTrimRight(cWD, "/") + "/"
		if _, err = os.Stat(cWD + exe); err == nil {
			failOnErr("%v", os.Chdir(cWD))
			cWD, err = os.Getwd()
			failOnErr("%v", err)
			return cWD, oriWD
		}
	}
	failPnOnErr(2, "%v: %s", eg.FILE_NOT_FOUND, exe)
	return "", ""
}
