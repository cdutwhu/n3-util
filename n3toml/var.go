package n3toml

import (
	"fmt"
	"strings"

	cmn "github.com/cdutwhu/json-util/common"
)

var (
	fPln        = fmt.Println
	sJoin       = strings.Join
	sSplit      = strings.Split
	sHasPrefix  = strings.HasPrefix
	sTrim       = strings.Trim
	sTrimLeft   = strings.TrimLeft
	sTrimRight  = strings.TrimRight
	sContains   = strings.Contains
	sReplaceAll = strings.ReplaceAll

	failOnErr      = cmn.FailOnErr
	rmTailFromLast = cmn.RmTailFromLast
	mustWriteFile  = cmn.MustWriteFile
)
