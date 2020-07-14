package n3xml

import (
	"io/ioutil"
	"os"
	"regexp"

	eg "github.com/cdutwhu/n3-util/n3errs"
	"github.com/go-xmlfmt/xmlfmt"
)

// SmashSave :
func SmashSave(xml, saveDir string) []string {
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		failOnErr("%v", os.MkdirAll(saveDir, os.ModePerm))
	}
	saveDir = sTrimRight(saveDir, `/\`) + "/"
	mObjCnt := make(map[string]int)

	SubRoots, Subs, err := Smash(RmComment(xml))
	failOnErr("%v", err)
	for i, subRoot := range SubRoots {
		filename := fSf("%s%s_%d.xml", saveDir, subRoot, mObjCnt[subRoot])
		// fPln(filename)

		Subs[i] = xmlfmt.FormatXML(Subs[i], "", "    ")
		Subs[i] = sTrim(Subs[i], " \n\r\t")
		// fPln(Subs[i])

		failOnErr("%v", ioutil.WriteFile(filename, []byte(Subs[i]), os.ModePerm))
		mObjCnt[subRoot]++
	}
	return Subs
}

// Smash :
func Smash(xml string) (SubRoots, Subs []string, err error) {
	if !isXML(xml) {
		return nil, nil, eg.PARAM_INVALID_XML
	}

	root := XMLRoot(xml)
	offset := len(fSf("<%s>", root)) + 1
	remain := xml[offset:]
	r := regexp.MustCompile(`<[^> /]+[ >]`)

	// I := 1

AGAIN:
	if start := r.FindString(remain); start != "" {
		subroot := sTrim(start, "<> \n\t\r")
		// fPln(I, subroot)

		SubRoots = append(SubRoots, subroot)

		endMark := fSf("</%s>", subroot)
		endPos := sIndex(remain, endMark)
		length := endPos + len(endMark)
		offset += length

		sub := remain[:length]
		failOnErrWhen(!isXML(sub), "%v", eg.XML_INVALID)
		Subs = append(Subs, sub)

		remain = xml[offset:]

		// I++
		goto AGAIN
	}

	return SubRoots, Subs, nil
}
