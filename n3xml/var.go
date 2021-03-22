package n3xml

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/misc"
	"github.com/cdutwhu/gotil/str"
)

var (
	fPln            = fmt.Println
	fSf             = fmt.Sprintf
	sTrim           = strings.Trim
	sTrimRight      = strings.TrimRight
	sIndex          = strings.Index
	sHasPrefix      = strings.HasPrefix
	sSplit          = strings.Split
	sReplaceAll     = strings.ReplaceAll
	rxMustCompile   = regexp.MustCompile
	failOnErr       = fn.FailOnErr
	failP1OnErrWhen = fn.FailP1OnErrWhen
	isXML           = judge.IsXML
	isNumeric       = judge.IsNumeric
	splitLn         = str.SplitLn
	trackTime       = misc.TrackTime
	mustWriteFile   = io.MustWriteFile
)
