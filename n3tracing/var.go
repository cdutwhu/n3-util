package n3tracing

import (
	"fmt"

	"github.com/cdutwhu/debog/fn"
)

var (
	fPln             = fmt.Println
	fEf              = fmt.Errorf
	failOnErr        = fn.FailOnErr
	failOnErrWhen    = fn.FailOnErrWhen
	failP1OnErrWhen  = fn.FailP1OnErrWhen
	logger           = fn.Logger
	loggerWhen       = fn.LoggerWhen
	warnOnErrWhen    = fn.WarnOnErrWhen
	enableWarnDetail = fn.EnableWarnDetail
)
