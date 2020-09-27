package n3xml

import (
	"os"
	"strings"

	"github.com/cdutwhu/gotil/iter"
)

// Break :
func Break(xml, saveDir string, fmteach bool) ([]string, bool) {
	if _, err := os.Stat(saveDir); saveDir != "" && os.IsNotExist(err) {
		failOnErr("%v", os.MkdirAll(saveDir, os.ModePerm))
	}
	if _, cont, _, _ := TagContAttrVal(xml); cont != "" {
		saveDir = sTrimRight(saveDir, `/\`) + "/"
		mObjCnt := make(map[string]int)
		// roots, subs := SmashCont(RmComment(cont))
		roots, subs := BreakCont(cont)
		for i, subRoot := range roots {
			if saveDir != "/" {
				filename := fSf("%s%s_%d.xml", saveDir, subRoot, mObjCnt[subRoot])
				if fmteach {
					subs[i] = Fmt(subs[i])
				}
				mustWriteFile(filename, []byte(subs[i]))
			}
			mObjCnt[subRoot]++
		}
		return subs, true
	}
	return nil, false
}

// BreakCont :
func BreakCont(xml string) (roots, subs []string) {

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

// Fast Process for Formatted XML

func mkProbe(ls ...int) (probes []string) {
	for _, l := range ls {
		var sb strings.Builder
		for i := 0; i < l; i++ {
			sb.WriteString("\t")
		}
		sb.WriteString("<")
		probes = append(probes, sb.String())
	}
	return
}

// FastBreakCont : xml is formatted.
func FastBreakCont(xml string) (roots, subs []string) {
	var sb strings.Builder
	probes := mkProbe(iter.Iter2Slc(0, 11)...)

	for _, ln := range sSplit(xml, "\n") {
		if sHasPrefix(ln, probes[1]) {
			if sb.Len() > 0 {
				subs = append(subs, sb.String())
			}
			sb.Reset()

			s, e := 0, 0
			for j, c := range ln {
				if c == '<' {
					s = j + 1
				}
				if c == '>' || c == ' ' || c == '/' {
					e = j
					break
				}
			}
			roots = append(roots, ln[s:e])
		}
		if sTrim(ln, " \t\r\n") != "" {
			_, err := sb.WriteString(ln)
			failOnErr("%v", err)
			_, err = sb.WriteString("\n")
			failOnErr("%v", err)
		}
	}
	return
}
