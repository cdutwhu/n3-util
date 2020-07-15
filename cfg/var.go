package cfg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/net"
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
	vof         = reflect.ValueOf
	typPTR      = reflect.Ptr
	typSTRUCT   = reflect.Struct

	isXML           = judge.IsXML
	failOnErr       = fn.FailOnErr
	failP1OnErr     = fn.FailP1OnErr
	failP1OnErrWhen = fn.FailP1OnErrWhen
	localIP         = net.LocalIP
)
