package n3toml

import (
	"fmt"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/str"
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

	failOnErr      = fn.FailOnErr
	rmTailFromLast = str.RmTailFromLast
	mustWriteFile  = io.MustWriteFile
)
