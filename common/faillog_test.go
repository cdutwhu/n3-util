package common

import "testing"

func TestFailLog(t *testing.T) {
	logfile := "./error.log"
	SetLog(logfile)
	defer ResetLog()

	msg := Log("hello")
	fPt(msg)

	LogWhen(1 < 3, "hello when")

	WarnOnErr("aa %v", fEf(""))
	FailOnErr("AA %v", nil)
	// FailOnErrWhen(1 > 0, "AA %v", fEf(""))
	// ResetLog()

	// logfile = "./log1.log"
	// SetLog(logfile)

	// logfile = "./log1.log"
	// SetLog(logfile)
	// // FailOnErr("%v", fEf("test panic"))

	// if e := WarnOnErrWhen(1 < 2, "%v", fEf("test")); e != nil {
	// 	fPln(e.Error())
	// }
}
