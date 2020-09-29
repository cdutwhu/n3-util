package n3xml

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/misc"
	"github.com/cdutwhu/gotil/rflx"
	"github.com/cdutwhu/gotil/str"
)

var (
	fEf             = fmt.Errorf
	fPln            = fmt.Println
	fPf             = fmt.Printf
	fSf             = fmt.Sprintf
	sTrim           = strings.Trim
	sTrimLeft       = strings.TrimLeft
	sTrimRight      = strings.TrimRight
	sIndex          = strings.Index
	sLastIndex      = strings.LastIndex
	sHasPrefix      = strings.HasPrefix
	sHasSuffix      = strings.HasSuffix
	sSplit          = strings.Split
	sCount          = strings.Count
	sReplace        = strings.Replace
	sReplaceAll     = strings.ReplaceAll
	rxMustCompile   = regexp.MustCompile
	failOnErr       = fn.FailOnErr
	failOnErrWhen   = fn.FailOnErrWhen
	failP1OnErrWhen = fn.FailP1OnErrWhen
	isXML           = judge.IsXML
	isNumeric       = judge.IsNumeric
	exist           = judge.Exist
	replByPosGrp    = str.ReplByPosGrp
	hasAnyPrefix    = str.HasAnyPrefix
	splitLn         = str.SplitLn
	trackTime       = misc.TrackTime
	toGeneralSlc    = rflx.ToGeneralSlc
	mustWriteFile   = io.MustWriteFile
)
