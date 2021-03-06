package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"time"

	eg "github.com/cdutwhu/n3-util/n3errs"
)

const tmFmt = "2006/01/02 15:04:05"

var (
	log2file                      = false
	mPathFile map[string]*os.File = make(map[string]*os.File)
)

// FetchLog2CSV :
func FetchLog2CSV(logFile, logType string, tmBackwards, tmOffset int, desc bool) (string, error) {
	logs, err := FetchLog(logFile, logType, tmBackwards, tmOffset, desc)
	if err != nil {
		return "", err
	}
	if len(logs) == 0 {
		return "", nil
	}
	content := "Time,Type,Desc\n"
	content += sReplaceAll(sJoin(logs, "\n"), " \t", ",")
	file := RmTailFromLast(logFile, ".") + "-" + logType + ".csv"
	MustWriteFile(file, []byte(content))
	return file, nil
}

// FetchLog2File :
func FetchLog2File(logFile, logType string, tmBackwards, tmOffset int, desc bool) (string, error) {
	logs, err := FetchLog(logFile, logType, tmBackwards, tmOffset, desc)
	if err != nil {
		return "", err
	}
	if len(logs) == 0 {
		return "", nil
	}
	content := sJoin(logs, "\n")
	file := RmTailFromLast(logFile, ".") + "-" + logType + "." + RmHeadToLast(logFile, ".")
	MustWriteFile(file, []byte(content))
	return file, nil
}

// FetchLog : logType [INFO, WARN, FAIL]; tmBackwards second unit
func FetchLog(logFile, logType string, tmBackwards, tmOffset int, desc bool) ([]string, error) {
	logTypes := []string{"INFO", "WARN", "FAIL"}
	if ok, err := XIn(logType, logTypes); err != nil || !ok {
		return nil, eg.PARAM_NOT_SUPPORTED
	}

	bytes, err := ioutil.ReadFile(logFile)
	// FailOnErr("%v", err)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	// zone, offset := now.Zone()
	// fPln(now)

	past := now.Add(-time.Second * time.Duration(tmBackwards))
	// fPln(past)

	re := regexp.MustCompile(fSf(`^[0-9/: ]{20}\t%s \t`, logType))
	logs := []string{}
	for _, ln := range sSplit(string(bytes), "\n") {
		if re.MatchString(ln) {
			tm, err := time.Parse(tmFmt, ln[:19])
			FailOnErr("%v", err)
			if tm.After(past) {
				tm = tm.Add(time.Second * time.Duration(tmOffset))
				ln = tm.Format(tmFmt) + ln[19:]
				logs = append(logs, ln)
			}
		}
	}

	if desc {
		for l, r := 0, len(logs)-1; l < r; l, r = l+1, r-1 {
			logs[l], logs[r] = logs[r], logs[l]
		}
	}

	return logs, nil
}

// Caller :
func Caller(fullpath bool) string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	if fullpath {
		return frame.Function
	}
	return RmHeadToLast(frame.Function, ".")
}

// FuncTrack : full path of func name
func FuncTrack(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// trackCaller :
func trackCaller() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc) // 3 is for util-FailLog. 2 is for "trackCaller" caller
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fSf("\n%s:%d\n%s\n", frame.File, frame.Line, frame.Function)
}

// IsFLog : is logging into local file
func IsFLog() bool {
	return log2file
}

// SetLog :
func SetLog(logpath string) {
	if abspath, err := filepath.Abs(logpath); err == nil {
		if f, ok := mPathFile[abspath]; ok {
			log.SetFlags(log.LstdFlags | log.LUTC)
			log.SetOutput(f)
			log2file = true
			return
		}
		if _, err := os.Stat(abspath); err == nil || os.IsNotExist(err) {
			if f, err := os.OpenFile(abspath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err == nil {
				mPathFile[abspath] = f
				log.SetFlags(log.LstdFlags | log.LUTC)
				log.SetOutput(f)
				log2file = true
			}
		}
	}
}

// ResetLog : call once at the exit
func ResetLog() {
	for logPath, f := range mPathFile {
		// delete empty error log
		fi, err := f.Stat()
		FailOnErr("%v", err)
		if fi.Size() == 0 {
			FailOnErr("%v", os.Remove(logPath))
		}
		// close
		f.Close()
	}
	mPathFile = make(map[string]*os.File)
	log.SetOutput(os.Stdout)
	log2file = false
}

// FailOnErr : error holder use "%v"
func FailOnErr(format string, v ...interface{}) {
	for _, p := range v {
		switch p.(type) {
		case error:
			{
				if p != nil {
					fatalInfo := fSf("\tFAIL \t\""+format+"\"%s\n", append(v, trackCaller())...)
					if log2file {
						fPln(time.Now().Format(tmFmt) + fatalInfo)
					}
					log.Fatalf(fatalInfo)
				}
			}
		}
	}
}

// FailOnErrWhen :
func FailOnErrWhen(condition bool, format string, v ...interface{}) {
	if condition {
		for _, p := range v {
			switch p.(type) {
			case error:
				{
					if p != nil {
						fatalInfo := fSf("\tFAIL \t\""+format+"\"%s\n", append(v, trackCaller())...)
						if log2file {
							fPln(time.Now().Format(tmFmt) + fatalInfo)
						}
						log.Fatalf(fatalInfo)
					}
				}
			}
		}
	}
}

// Log :
func Log(format string, v ...interface{}) string {
	tc := trackCaller()
	logItem := fSf("\tINFO \t\""+format+"\"%s\n", append(v, tc)...)
	log.Printf("%s", logItem)
	return time.Now().Format(tmFmt) + RmTailFromLast(logItem, tc)
}

// LogWhen :
func LogWhen(condition bool, format string, v ...interface{}) string {
	if condition {
		tc := trackCaller()
		logItem := fSf("\tINFO \t\""+format+"\"%s\n", append(v, tc)...)
		log.Printf("%s", logItem)
		return time.Now().Format(tmFmt) + RmTailFromLast(logItem, tc)
	}
	return ""
}

// WarnOnErr :
func WarnOnErr(format string, v ...interface{}) error {
	for _, p := range v {
		switch p.(type) {
		case error:
			{
				if p != nil {
					tc := trackCaller()
					warnItem := fSf("\tWARN \t\""+format+"\"%s\n", append(v, tc)...)
					log.Printf(warnItem)
					return fEf("%v", time.Now().Format(tmFmt)+RmTailFromLast(warnItem, tc))
				}
			}
		}
	}
	return nil
}

// WarnOnErrWhen :
func WarnOnErrWhen(condition bool, format string, v ...interface{}) error {
	if condition {
		for _, p := range v {
			switch p.(type) {
			case error:
				{
					if p != nil {
						tc := trackCaller()
						warnItem := fSf("\tWARN \t\""+format+"\"%s\n", append(v, tc)...)
						log.Printf(warnItem)
						return fEf("%v", time.Now().Format(tmFmt)+RmTailFromLast(warnItem, tc))
					}
				}
			}
		}
	}
	return nil
}

// WrapOnErr :
func WrapOnErr(format string, v ...interface{}) error {
	for _, p := range v {
		switch p.(type) {
		case error:
			{
				if p != nil {
					return fmt.Errorf(format, v...)
				}
			}
		}
	}
	return nil
}
