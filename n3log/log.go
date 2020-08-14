package n3log

import (
	"bytes"
	"flag"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// lrInit :
func lrInit() {
	profile := flag.String("lrprofile", "test", "logrus formatter")
	flag.Parse()
	if *profile == "dev" {
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000",
			FullTimestamp:   true,
		})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}

// SetLoggly :
func SetLoggly(enable bool, token, tag string) {
	loggly = enable
	urlR = sReplace(urlT, "#token#", token, 1)
	urlR = sReplace(urlR, "#tag#", tag, 1)
	if loggly {
		lrInit()
	}
}

// Loggly :
func Loggly(level string) Output {
	return func(format string, args ...interface{}) {
		var logwriter bytes.Buffer
		logrus.SetOutput(io.Writer(&logwriter))
		fn := logrus.Printf
		switch level {
		case "info":
			fn = logrus.Printf
		case "warn", "warning":
			fn = logrus.Warnf
		case "error":
			fn = logrus.Errorf
		case "debug":
			fn = logrus.Debugf
		default:
			fn = logrus.Printf
		}
		fn(format, args...)
		if loggly {
			_, err := http.Post(urlR, "application/json", io.Reader(&logwriter))
			warnOnErr("%v", err)
		}
		fPt(logwriter.String())
	}
}

// --------------------------------------------------------- //

// Output :
type Output func(format string, args ...interface{})

// Logger : group of Output
type Logger []Output

// Bind :
func Bind(outputs ...Output) Logger {
	return append([]Output{}, outputs...)
}

// Do :
func (lg Logger) Do(format string, args ...interface{}) {
	// defer trackTime(time.Now())
	for _, f := range lg {
		go f(format, args...)
	}
}
