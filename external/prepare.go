package external

import (
	"os"

	"github.com/cdutwhu/n3-util/n3err"
)

// Prepare : // append `defer func() { os.Chdir(oriWD) }()` after this function
func Prepare(exe string, wkDirs ...string) (cWD, oriWD string) {
	oriWD, err := os.Getwd()
	failOnErr("%v", err)
	for _, cWD = range append(wkDirs, "./", "../", "../../", "/usr/bin") {
		cWD = sTrimRight(cWD, "/") + "/"
		if _, err = os.Stat(cWD + exe); err == nil {
			failOnErr("%v", os.Chdir(cWD))
			cWD, err = os.Getwd()
			failOnErr("%v", err)
			return cWD, oriWD
		}
	}

	panic(fSf("%v: %s", n3err.FILE_NOT_FOUND, exe))
	// failPnOnErr(2, "%v: %s", n3err.FILE_NOT_FOUND, exe)
	// return "", ""
}
