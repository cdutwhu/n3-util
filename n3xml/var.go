package n3xml

import (
	"fmt"
	"strings"

	cmn "github.com/cdutwhu/n3-util/common"
)

var (
	fEf        = fmt.Errorf
	fPln       = fmt.Println
	fSf        = fmt.Sprintf
	sTrim      = strings.Trim
	sTrimRight = strings.TrimRight
	sIndex     = strings.Index

	failOnErr     = cmn.FailOnErr
	failOnErrWhen = cmn.FailOnErrWhen
	isXML         = cmn.IsXML
	xmlRoot       = cmn.XMLRoot
	replByPosGrp  = cmn.ReplByPosGrp
)
