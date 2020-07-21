package n3log

import (
	"bytes"
	"flag"
	"io"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	urlT   = `http://logs-01.loggly.com/inputs/#token#/tag/http/`
	urlR   = ""
	loggly = false
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

// LrOut :
func LrOut(logrusOut func(format string, args ...interface{}), format string, args ...interface{}) (msg string, err error) {
	var logwriter bytes.Buffer
	logrus.SetOutput(io.Writer(&logwriter))
	logrusOut(format, args...)
	if loggly {
		_, err = http.Post(urlR, "application/json", io.Reader(&logwriter))
	}
	return logwriter.String(), err
}

// SetLogglyToken :
func SetLogglyToken(token string) {
	urlR = strings.Replace(urlT, "#token#", token, 1)
}

// EnableLoggly :
func EnableLoggly(enable bool) {
	loggly = enable
}
