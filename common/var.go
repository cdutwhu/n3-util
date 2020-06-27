package common

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	fPt         = fmt.Print
	fPf         = fmt.Printf
	fPln        = fmt.Println
	fSf         = fmt.Sprintf
	fSp         = fmt.Sprint
	fEf         = fmt.Errorf
	sSplit      = strings.Split
	sJoin       = strings.Join
	sCount      = strings.Count
	sReplace    = strings.Replace
	sReplaceAll = strings.ReplaceAll
	sIndex      = strings.Index
	sLastIndex  = strings.LastIndex
	sTrim       = strings.Trim
	sTrimLeft   = strings.TrimLeft
	sHasPrefix  = strings.HasPrefix
	sHasSuffix  = strings.HasSuffix
	sContains   = strings.Contains
	sToUpper    = strings.ToUpper
	scParseUint = strconv.ParseUint
)

var (
	repParam = regexp.MustCompile(`^\$[0-9]+$`)
)
