package common

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	log2file                      = false
	mPathFile map[string]*os.File = make(map[string]*os.File)
)

// FuncTrack : full path of func name
// func FuncTrack(i interface{}) string {
// 	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
// }

// trackCaller :
func trackCaller() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc) // 3 is for util-FailLog. 2 is for "trackCaller" caller
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fSf("\n%s:%d\n%s\n", frame.File, frame.Line, frame.Function)
}

// SetLog :
func SetLog(logpath string) {
	if abspath, err := filepath.Abs(logpath); err == nil {
		if f, ok := mPathFile[abspath]; ok {
			log.SetFlags(log.LstdFlags)
			log.SetOutput(f)
			log2file = true
			return
		}
		if _, err := os.Stat(abspath); err == nil || os.IsNotExist(err) {
			if f, err := os.OpenFile(abspath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err == nil {
				mPathFile[abspath] = f
				log.SetFlags(log.LstdFlags)
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
					fatalInfo := fSf("FAIL: "+format+"%s\n", append(v, trackCaller())...)
					if log2file {
						fPln(time.Now().Format("2006/01/02 15:04:05 ") + fatalInfo)
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
						fatalInfo := fSf("FAIL: "+format+"%s\n", append(v, trackCaller())...)
						if log2file {
							fPln(time.Now().Format("2006/01/02 15:04:05 ") + fatalInfo)
						}
						log.Fatalf(fatalInfo)
					}
				}
			}
		}
	}
}

// Log :
func Log(format string, v ...interface{}) (logItem string) {
	logItem = fSf("INFO: "+format+"%s\n", append(v, trackCaller())...)
	log.Printf("%s", logItem)
	if log2file {
		logItem = time.Now().Format("2006/01/02 15:04:05 ") + logItem
	}
	return
}

// LogWhen :
func LogWhen(condition bool, format string, v ...interface{}) (logItem string) {
	if condition {
		logItem = fSf("INFO: "+format+"%s\n", append(v, trackCaller())...)
		log.Printf("%s", logItem)
		if log2file {
			logItem = time.Now().Format("2006/01/02 15:04:05 ") + logItem
		}
	}
	return
}

// WarnOnErr :
func WarnOnErr(format string, v ...interface{}) error {
	for _, p := range v {
		switch p.(type) {
		case error:
			{
				if p != nil {
					log.Printf("WARN: "+format+"%s\n", append(v, trackCaller())...)
					return p.(error)
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
						log.Printf("WARN: "+format+"%s\n", append(v, trackCaller())...)
						return p.(error)
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
