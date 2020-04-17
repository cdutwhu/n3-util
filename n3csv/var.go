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
	sReplaceAll = strings.ReplaceAll

	failOnErr      = cmn.FailOnErr
	failOnErrWhen  = cmn.FailOnErrWhen
	mustWriteFile  = cmn.MustWriteFile
	JSONScalarSelX = n3json.ScalarSelX
)
