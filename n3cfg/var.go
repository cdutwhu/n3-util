package n3cfg

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"sync"

	"github.com/cdutwhu/debog/base"
	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/net"
	"github.com/cdutwhu/gotil/rflx"
	"github.com/cdutwhu/gotil/str"
	"github.com/cdutwhu/n3-util/external"
)

var (
	fSf             = fmt.Sprintf
	fPln            = fmt.Println
	fEf             = fmt.Errorf
	sSplit          = strings.Split
	sHasPrefix      = strings.HasPrefix
	sHasSuffix      = strings.HasSuffix
	sTrim           = strings.Trim
	sContains       = strings.Contains
	sCount          = strings.Count
	sReplaceAll     = strings.ReplaceAll
	rMustCompile    = regexp.MustCompile
	vof             = reflect.ValueOf
	typPTR          = reflect.Ptr
	typSTRUCT       = reflect.Struct
	callerSrc       = base.CallerSrc
	trackCaller     = base.TrackCaller
	failOnErr       = fn.FailOnErr
	failP1OnErr     = fn.FailP1OnErr
	failP1OnErrWhen = fn.FailP1OnErrWhen
	failPnOnErr     = fn.FailPnOnErr
	isNumeric       = judge.IsNumeric
	localIP         = net.LocalIP
	prepare         = external.Prepare
	struct2Env      = rflx.Struct2Env
	splitLn         = str.SplitLn
	rmTailFromLast  = str.RmTailFromLast
)

var (
	mux sync.Mutex
)
