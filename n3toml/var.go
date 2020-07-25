package n3toml

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/str"
)

var (
	fSln        = fmt.Sprintln
	fSf         = fmt.Sprintf
	fPln        = fmt.Println
	fPf         = fmt.Printf
	sIndex      = strings.Index
	sJoin       = strings.Join
	sSplit      = strings.Split
	sHasPrefix  = strings.HasPrefix
	sHasSuffix  = strings.HasSuffix
	sTrim       = strings.Trim
	sTrimLeft   = strings.TrimLeft
	sTrimRight  = strings.TrimRight
	sContains   = strings.Contains
	sCount      = strings.Count
	sReplaceAll = strings.ReplaceAll
	sToUpper    = strings.ToUpper
	ucIsUpper   = unicode.IsUpper

	failOnErr       = fn.FailOnErr
	failP1OnErr     = fn.FailP1OnErr
	failP1OnErrWhen = fn.FailP1OnErrWhen
	rmTailFromLast  = str.RmTailFromLast
	rmTailFromFirst = str.RmTailFromFirst
	rmHeadToFirst   = str.RmHeadToFirst
	hasAnyPrefix    = str.HasAnyPrefix
	mustWriteFile   = io.MustWriteFile
	isNumeric       = judge.IsNumeric
)
