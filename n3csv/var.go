package n3csv

import (
	"fmt"
	"strings"

	cmn "github.com/cdutwhu/json-util/common"
	n3json "github.com/cdutwhu/json-util/n3json"
)

var (
	fPln        = fmt.Println
	fEf         = fmt.Errorf
	fSf         = fmt.Sprintf
	sReplaceAll = strings.ReplaceAll

	failOnErr     = cmn.FailOnErr
	failOnErrWhen = cmn.FailOnErrWhen
	warnOnErr     = cmn.WarnOnErr
	mustWriteFile = cmn.MustWriteFile
	setLog        = cmn.SetLog
	resetLog      = cmn.ResetLog

	jsonScalarSelX = n3json.ScalarSelX
)
