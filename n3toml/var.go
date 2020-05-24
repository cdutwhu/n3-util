package n3toml

import (
	"fmt"
	"strings"

	cmn "github.com/cdutwhu/n3-util/common"
)

var (
	fPln        = fmt.Println
	sJoin       = strings.Join
	sSplit      = strings.Split
	sHasPrefix  = strings.HasPrefix
	sHasSuffix  = strings.HasSuffix
	sTrim       = strings.Trim
	sTrimLeft   = strings.TrimLeft
	sTrimRight  = strings.TrimRight
	sContains   = strings.Contains
	sReplaceAll = strings.ReplaceAll

	failOnErr      = cmn.FailOnErr
	rmTailFromLast = cmn.RmTailFromLast
	mustWriteFile  = cmn.MustWriteFile
)
