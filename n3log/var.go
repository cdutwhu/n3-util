package n3log

import (
	"fmt"
	"log"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/misc"
)

var (
	fSf           = fmt.Sprintf
	fPln          = fmt.Println
	fPt           = fmt.Print
	lPf           = log.Printf
	sReplace      = strings.Replace
	logger        = fn.Logger
	warnOnErr     = fn.WarnOnErr
	warnOnErrWhen = fn.WarnOnErrWhen
	trackTime     = misc.TrackTime
)

var (
	urlT    = `http://logs-01.loggly.com/inputs/#token#/tag/#tag#/`
	urlR    = ""
	loggly  = false
	syncLog = false
)
