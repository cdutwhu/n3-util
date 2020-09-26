package n3xml

import (
	"os"
)

// SmashFirstAndSave :
func SmashFirstAndSave(xml, saveDir string) ([]string, bool) {
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		failOnErr("%v", os.MkdirAll(saveDir, os.ModePerm))
	}
	saveDir = sTrimRight(saveDir, `/\`) + "/"
	mObjCnt := make(map[string]int)

	if _, cont, _, _ := TagContAttrVal(xml); cont != "" {
		roots, subs := SmashCont(RmComment(cont))
		for i, subRoot := range roots {
			filename := fSf("%s%s_%d.xml", saveDir, subRoot, mObjCnt[subRoot])
			subs[i] = Fmt(subs[i])
			mustWriteFile(filename, []byte(subs[i]))
			mObjCnt[subRoot]++
		}
		return subs, true
	}
	return nil, false
}

// SmashCont :
func SmashCont(xml string) (roots, subs []string) {

	remain := xml
AGAIN:
	if loc := rxTag.FindStringIndex(remain); loc != nil {
		s, e := loc[0], loc[1]
		root := remain[s+1 : e-1]
		roots = append(roots, root)

		remain = remain[s:] // from first '<tag>'
		// fPln("remain:", remain)

		end1, end2 := -1, -1
		if loc := rxMustCompile(fSf(`</%s\s*>`, root)).FindStringIndex(remain); loc != nil {
			_, end1 = loc[0], loc[1] // update e to '</tag>' end
			// fPln("end:", remain[s:end1]) // end tag
		}
		if i := sIndex(remain, "/>"); i >= 0 {
			end2 = i + 2 // update e to '/>' end
		}

		// if '/>' is found, and before '</tag>', and this part is valid XML
		switch {
		case end1 >= 0 && end2 < 0:
			e = end1
		case end1 < 0 && end2 >= 0:
			e = end2
		case end1 >= 0 && end2 >= 0:
			if end2 < end1 && isXML(remain[:end2]) {
				e = end2
			} else {
				e = end1
			}
		default:
			panic("invalid sub xml")
		}

		sub := remain[:e]
		// fPln("sub:", sub)
		subs = append(subs, sub)

		remain = remain[e:] // from end of first '</tag>' or '/>'
		goto AGAIN
	}

	return
}
