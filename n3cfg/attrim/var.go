package attrim

import (
	"fmt"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/str"
)

var (
	fPln            = fmt.Println
	sJoin           = strings.Join
	sSplit          = strings.Split
	sHasPrefix      = strings.HasPrefix
	sHasSuffix      = strings.HasSuffix
	sTrim           = strings.Trim
	sTrimLeft       = strings.TrimLeft
	sTrimRight      = strings.TrimRight
	sContains       = strings.Contains
	sReplaceAll     = strings.ReplaceAll
	failOnErr       = fn.FailOnErr
	failP1OnErr     = fn.FailP1OnErr
	failP1OnErrWhen = fn.FailP1OnErrWhen
	failPnOnErr     = fn.FailPnOnErr
	rmTailFromLast  = str.RmTailFromLast
	rmTailFromFirst = str.RmTailFromFirst
	rmHeadToFirst   = str.RmHeadToFirst
	mustWriteFile   = io.MustWriteFile
)
