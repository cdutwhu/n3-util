package common

import (
	"io/ioutil"
	"testing"
	"time"

	eg "github.com/cdutwhu/n3-util/n3errs"
)

type fighter struct {
	name string
}

func (p *fighter) showName() {
	fPln(p.name + " @ " + Caller(false))
}

func TestFuncTrack(t *testing.T) {
	fPln(Caller(true))
	fPln(Caller(false))
	p := &fighter{name: "HAO HAIDONG"}
	p.showName()
}

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

func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func TestFetchLog(t *testing.T) {

	for _, name := range []string{
		"",
		"Local",
		"Asia/Shanghai",
		"America/New_York",
		"Australia/Melbourne",
	} {
		t, err := TimeIn(time.Now(), name)
		if err == nil {
			fPln(t.Location(), t.Format("15:04"))
		} else {
			fPln(name, "<time unknown>")
		}
	}

	fPln(" --------------------------------------- ")

	now := time.Now()
	zone, offset := now.Zone()
	fPln(zone, offset)

	fPln(" --------------------------------------- ")

	logs, err := FetchLog("./error.log", "WARN", 10000, 36000, true)
	FailOnErr("%v", err)
	for _, ln := range logs {
		fPln(ln)
	}
	FetchLog2File("./error.log", "FAIL", 10000, 36000, true)
	FetchLog2CSV("./error.log", "FAIL", 10000, 36000, true)
}

func ReadFile(path string) {
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.Name() != sToUpper(f.Name()[:1])+f.Name()[1:] {
			continue
		}
		if f.IsDir() {
			ReadFile(path + "/" + f.Name())
		} else {
			fPln((path + "/" + f.Name())[1:])
		}
	}
}

func TestListAllLoc(t *testing.T) {

	for _, zoneDir := range []string{
		// Update path according to your OS
		"/usr/share/zoneinfo/",
		"/usr/share/lib/zoneinfo/",
		"/usr/lib/locale/TZ/",
	} {
		ReadFile(zoneDir)
	}
}
