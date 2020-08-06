package strugen

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/cdutwhu/debog/base"
	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/str"
)

var (
	fSln            = fmt.Sprintln
	fSf             = fmt.Sprintf
	fPln            = fmt.Println
	sSplit          = strings.Split
	sHasPrefix      = strings.HasPrefix
	sHasSuffix      = strings.HasSuffix
	sTrim           = strings.Trim
	sTrimLeft       = strings.TrimLeft
	sTrimRight      = strings.TrimRight
	sContains       = strings.Contains
	sCount          = strings.Count
	sReplaceAll     = strings.ReplaceAll
	sToUpper        = strings.ToUpper
	sToLower        = strings.ToLower
	sTitle          = strings.Title
	ucIsUpper       = unicode.IsUpper
	failOnErr       = fn.FailOnErr
	failP1OnErr     = fn.FailP1OnErr
	failP1OnErrWhen = fn.FailP1OnErrWhen
	failPnOnErr     = fn.FailPnOnErr
	logger          = fn.Logger
	enableLog2F     = fn.EnableLog2F
	rmTailFromLast  = str.RmTailFromLast
	rmTailFromFirst = str.RmTailFromFirst
	rmHeadToFirst   = str.RmHeadToFirst
	hasAnyPrefix    = str.HasAnyPrefix
	mustWriteFile   = io.MustWriteFile
	isNumeric       = judge.IsNumeric
	isXML           = judge.IsXML
	callerSrc       = base.CallerSrc
)

const (
	logfile = "/home/qmiao/Desktop/debug.log"
)
