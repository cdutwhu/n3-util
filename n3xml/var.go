package n3xml

import (
	"fmt"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/str"
	"github.com/cdutwhu/n3-util/n3err"
)

var (
	fEf        = fmt.Errorf
	fPln       = fmt.Println
	fSf        = fmt.Sprintf
	sTrim      = strings.Trim
	sTrimRight = strings.TrimRight
	sIndex     = strings.Index

	failOnErr       = fn.FailOnErr
	failOnErrWhen   = fn.FailOnErrWhen
	failP1OnErrWhen = fn.FailP1OnErrWhen
	isXML           = judge.IsXML
	replByPosGrp    = str.ReplByPosGrp
)

// XMLRoot :
func XMLRoot(xml string) string {
	xml = sTrim(xml, " \t\n\r")
	failP1OnErrWhen(!isXML(xml), "%v", n3err.PARAM_INVALID_XML)

	start, end := 0, 0
	for i := len(xml) - 1; i >= 0; i-- {
		switch xml[i] {
		case '>':
			end = i
		case '/':
			start = i + 1
		}
		if start != 0 && end != 0 {
			break
		}
	}
	return xml[start:end]

	// check, flag (?s) let . includes "NewLine"
	// re1 := regexp.MustCompile(fSf(`(?s)^<%s .+</%s>$`, root, root))
	// re2 := regexp.MustCompile(fSf(`(?s)^<%s>.+</%s>$`, root, root))
	// failP1OnErrWhen(!re1.MatchString(xml) && !re2.MatchString(xml), "%v", n3err.XML_INVALID)
	// return
}
