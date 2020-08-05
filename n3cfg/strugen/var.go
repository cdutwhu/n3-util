package strugen

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/cdutwhu/debog/base"
	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/net"
	"github.com/cdutwhu/gotil/rflx"
	"github.com/cdutwhu/gotil/str"
	"github.com/cdutwhu/n3-util/external"
)

var (
	fSln         = fmt.Sprintln
	fSf          = fmt.Sprintf
	fPln         = fmt.Println
	fPf          = fmt.Printf
	fPt          = fmt.Print
	fSp          = fmt.Sprint
	fEf          = fmt.Errorf
	sIndex       = strings.Index
	sJoin        = strings.Join
	sSplit       = strings.Split
	sHasPrefix   = strings.HasPrefix
	sHasSuffix   = strings.HasSuffix
	sTrim        = strings.Trim
	sTrimLeft    = strings.TrimLeft
	sTrimRight   = strings.TrimRight
	sContains    = strings.Contains
	sCount       = strings.Count
	sReplaceAll  = strings.ReplaceAll
	sToUpper     = strings.ToUpper
	sReplace     = strings.Replace
	sLastIndex   = strings.LastIndex
	ucIsUpper    = unicode.IsUpper
	scParseUint  = strconv.ParseUint
	rMustCompile = regexp.MustCompile
	vof          = reflect.ValueOf
	typPTR       = reflect.Ptr
	typSTRUCT    = reflect.Struct

	failOnErr       = fn.FailOnErr
	failP1OnErr     = fn.FailP1OnErr
	failP1OnErrWhen = fn.FailP1OnErrWhen
	failPnOnErr     = fn.FailPnOnErr
	rmTailFromLast  = str.RmTailFromLast
	rmTailFromFirst = str.RmTailFromFirst
	rmHeadToFirst   = str.RmHeadToFirst
	hasAnyPrefix    = str.HasAnyPrefix
	mustWriteFile   = io.MustWriteFile
	isNumeric       = judge.IsNumeric
	isXML           = judge.IsXML
	localIP         = net.LocalIP
	prepare         = external.Prepare
	struct2Env      = rflx.Struct2Env
	env2Struct      = rflx.Env2Struct
	callerSrc       = base.CallerSrc
)
