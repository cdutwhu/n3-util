package n3log

import (
	"bytes"
	"flag"
	"io"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

// LrInit :
func LrInit() {
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
		fPln(logwriter.String())
	}
}

// SetLogglyToken :
func SetLogglyToken(token string) {
	urlR = strings.Replace(urlT, "#token#", token, 1)
}

// EnableLoggly :
func EnableLoggly(enable bool) {
	loggly = enable
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
	for _, f := range lg {
		f(format, args...)
	}
}
