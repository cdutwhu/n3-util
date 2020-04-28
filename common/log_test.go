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

	// FailOnErr("%v", eg.FOR_TEST)

	FailOnErrWhen(1 == 0, "%v", eg.FOR_TEST)
}

func TestExtractLog(t *testing.T) {
	logs := ExtractLog("./error.log", "WARN", 10000, 36000, true)
	for _, ln := range logs {
		fPln(ln)
	}
	ExtractLog2File("./error.log", "FAIL", 10000, 36000, true)
	ExtractLog2CSV("./error.log", "FAIL", 10000, 36000, true)
}
