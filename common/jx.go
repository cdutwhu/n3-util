package common

import (
	"encoding/json"
	"encoding/xml"

	eg "github.com/cdutwhu/n3-util/n3errs"
)

// IsXML :
func IsXML(str string) bool {
	return xml.Unmarshal([]byte(str), new(interface{})) == nil
}

// XMLRoot :
func XMLRoot(xml string) (root string) {
	xml = sTrim(xml, " \t\n\r")
	FailOnErrWhen(!IsXML(xml), "%v", eg.XML_INVALID)

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
	// FailOnErrWhen(!re1.MatchString(xml) && !re2.MatchString(xml), "%v", eg.XML_INVALID)
	// return
}

// IsJSON :
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// JSONRoot :
func JSONRoot(jsonstr string) string {
	x := make(map[string]interface{})
	FailOnErr("%v", json.Unmarshal([]byte(jsonstr), &x))
	for k := range x {
		return k
	}
	return ""
}
