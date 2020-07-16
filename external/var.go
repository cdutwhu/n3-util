package external

import (
	"fmt"
	"strings"

	"github.com/cdutwhu/debog/fn"
)

var (
	fPf = fmt.Printf
	fEf = fmt.Errorf
	fSf = fmt.Sprintf

	sReplaceAll = strings.ReplaceAll
	sHasPrefix  = strings.HasPrefix
	sHasSuffix  = strings.HasSuffix
	sTrimRight  = strings.TrimRight

	failOnErr   = fn.FailOnErr
	failP1OnErr = fn.FailP1OnErr
	failPnOnErr = fn.FailPnOnErr
)
