package n3xml

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/misc"
	"github.com/cdutwhu/gotil/str"
)

var (
	fEf             = fmt.Errorf
	fPln            = fmt.Println
	fSf             = fmt.Sprintf
	sTrim           = strings.Trim
	sTrimLeft       = strings.TrimLeft
	sTrimRight      = strings.TrimRight
	sIndex          = strings.Index
	sHasPrefix      = strings.HasPrefix
	sHasSuffix      = strings.HasSuffix
	sSplit          = strings.Split
	sCount          = strings.Count
	rMustCompile    = regexp.MustCompile
	failOnErr       = fn.FailOnErr
	failOnErrWhen   = fn.FailOnErrWhen
	failP1OnErrWhen = fn.FailP1OnErrWhen
	isXML           = judge.IsXML
	isNumeric       = judge.IsNumeric
	replByPosGrp    = str.ReplByPosGrp
	hasAnyPrefix    = str.HasAnyPrefix
	splitLn         = str.SplitLn
	trackTime       = misc.TrackTime
)
