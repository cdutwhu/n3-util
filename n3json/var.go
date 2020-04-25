package n3json

import (
	"fmt"
	"strings"

	cmn "github.com/cdutwhu/json-util/common"
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

	xin           = cmn.XIn
	isJSON        = cmn.IsJSON
	failOnErr     = cmn.FailOnErr
	failOnErrWhen = cmn.FailOnErrWhen
	replByPosGrp  = cmn.ReplByPosGrp
	trackTime     = cmn.TrackTime
	mustWriteFile = cmn.MustWriteFile
	indent        = cmn.Indent
	hasAnySuffix  = cmn.HasAnySuffix
)
