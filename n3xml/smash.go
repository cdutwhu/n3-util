package n3xml

import (
	"os"
	"regexp"

	"github.com/cdutwhu/n3-util/n3err"
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
		Subs[i] = Fmt(Subs[i])
		mustWriteFile(filename, []byte(Subs[i]))
		mObjCnt[subRoot]++
	}
	return Subs
}

// Smash :
func Smash(xml string) (SubRoots, Subs []string, err error) {
	if !isXML(xml) {
		return nil, nil, n3err.PARAM_INVALID_XML
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
		failOnErrWhen(!isXML(sub), "%v", n3err.XML_INVALID)
		Subs = append(Subs, sub)

		remain = xml[offset:]

		// I++
		goto AGAIN
	}

	return SubRoots, Subs, nil
}
