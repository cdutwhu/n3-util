package n3log

import (
	"fmt"
	"log"

	"github.com/cdutwhu/debog/fn"
)

var (
	fSf           = fmt.Sprintf
	fPln          = fmt.Println
	lPf           = log.Printf
	logger        = fn.Logger
	warnOnErr     = fn.WarnOnErr
	warnOnErrWhen = fn.WarnOnErrWhen
)

var (
	urlT   = `http://logs-01.loggly.com/inputs/#token#/tag/http/`
	urlR   = ""
	loggly = false
)
