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
	sIndex      = strings.Index
	sJoin       = strings.Join
	sSplit      = strings.Split
	sHasPrefix  = strings.HasPrefix
	sHasSuffix  = strings.HasSuffix
	sTrim       = strings.Trim
	sTrimLeft   = strings.TrimLeft
	sTrimRight  = strings.TrimRight
	sContains   = strings.Contains
	sReplaceAll = strings.ReplaceAll

	failOnErr       = fn.FailOnErr
	rmTailFromLast  = str.RmTailFromLast
	rmTailFromFirst = str.RmTailFromFirst
	mustWriteFile   = io.MustWriteFile
)
