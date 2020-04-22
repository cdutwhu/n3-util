package common

import (
	"testing"

	eg "github.com/cdutwhu/json-util/n3errs"
)

func TestFailLog(t *testing.T) {
	logfile := "./error.log"
	SetLog(logfile)
	defer ResetLog()

	if msg := Log("hello"); IsFLog() {
		fPln("***", msg)
	}

	if msg := LogWhen(1 < 3, "hello when"); IsFLog() {
		fPln("***", msg)
	}

	if e := WarnOnErr("%v: hello WarnOnErr", eg.FOR_TEST); e != nil && IsFLog() {
		fPf("*** %v\n", e)
	}

	FailOnErr("%v", nil)

	FailOnErrWhen(1 == 0, "%v", eg.FOR_TEST)
}
