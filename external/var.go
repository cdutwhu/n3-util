package external

import (
	"fmt"
	"strings"

	cmn "github.com/cdutwhu/n3-util/common"
)

var (
	fPf = fmt.Printf
	fEf = fmt.Errorf

	sReplaceAll = strings.ReplaceAll
	sHasPrefix  = strings.HasPrefix
	sHasSuffix  = strings.HasSuffix
	sTrimRight  = strings.TrimRight

	failOnErr = cmn.FailOnErr
)
