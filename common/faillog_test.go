package common

import (
	"testing"

	eg "github.com/cdutwhu/json-util/n3errs"
)

func TestFailLog(t *testing.T) {
	logfile := "./error.log"
	SetLog(logfile)
	defer ResetLog()

	msg := Log("hello")
	fPt(msg)

	LogWhen(1 < 3, "hello when")

	WarnOnErr("aa %v", eg.FOR_TEST)
	FailOnErr("AA %v", nil)
	// FailOnErrWhen(1 > 0, "AA %v", eg.FOR_TEST)
	// ResetLog()

	// logfile = "./log1.log"
	// SetLog(logfile)

	// logfile = "./log1.log"
	// SetLog(logfile)
	// // FailOnErr("%v", eg.FOR_TEST)

	// if e := WarnOnErrWhen(1 < 2, "%v", eg.FOR_TEST); e != nil {
	// 	fPln(e.Error())
	// }
}
