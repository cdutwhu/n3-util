package n3json

import (
	"fmt"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/misc"
	"github.com/cdutwhu/gotil/rflx"
	"github.com/cdutwhu/gotil/str"
)

// BLANK :
const BLANK = " \t\r\n"

var (
	fPf         = fmt.Printf
	fPln        = fmt.Println
	fSf         = fmt.Sprintf
	fEf         = fmt.Errorf
	sSplit      = strings.Split
	sJoin       = strings.Join
	sCount      = strings.Count
	sContains   = strings.Contains
	sReplace    = strings.Replace
	sReplaceAll = strings.ReplaceAll
	sIndex      = strings.Index
	sLastIndex  = strings.LastIndex
	sTrim       = strings.Trim
	sTrimLeft   = strings.TrimLeft
	sTrimRight  = strings.TrimRight
	sHasPrefix  = strings.HasPrefix
	sHasSuffix  = strings.HasSuffix
	sRepeat     = strings.Repeat

	exist           = judge.Exist
	isJSON          = judge.IsJSON
	mustWriteFile   = io.MustWriteFile
	failOnErr       = fn.FailOnErr
	failOnErrWhen   = fn.FailOnErrWhen
	failP1OnErr     = fn.FailP1OnErr
	failP1OnErrWhen = fn.FailP1OnErrWhen
	trackTime       = misc.TrackTime
	replByPosGrp    = str.ReplByPosGrp
	indent          = str.IndentTxt
	hasAnySuffix    = str.HasAnySuffix
	toGeneralSlc    = rflx.ToGeneralSlc
)
