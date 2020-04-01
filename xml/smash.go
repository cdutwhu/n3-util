package xml

import (
	"regexp"

	cmn "github.com/cdutwhu/json-util/common"
)

func smash(xml string) (lsSub []string) {
	cmn.FailOnErrWhen(!cmn.IsXML(xml), "%v", fEf("Invalid XML"))

	root := cmn.XMLRoot(xml)
	offset := len(fSf("<%s>", root)) + 1
	remain := xml[offset:]
	r := regexp.MustCompile(`<[^> /]+[ >]`)

	I := 1

AGAIN:
	if start := r.FindString(remain); start != "" {
		subroot := sTrim(start, "<> \n\t\r")
		fPln(I, subroot)

		endMark := fSf("</%s>", subroot)
		endPos := sIndex(remain, endMark)
		length := endPos + len(endMark)
		offset += length

		sub := remain[:length]
		// cmn.FailOnErrWhen(!cmn.IsXML(sub), "%v", fEf("Invalid XML"))
		lsSub = append(lsSub, sub)

		remain = xml[offset:]

		I++
		goto AGAIN
	}

	// fPln(lsSub[0])

	return lsSub
}
