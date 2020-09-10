package n3xml

import (
	"github.com/cdutwhu/n3-util/n3err"
	"github.com/go-xmlfmt/xmlfmt"
)

// Fmt :
func Fmt(xml string) string {
	xml = sReplaceAll(xml, "\r\n", "\n")                               // "\r\n" -> "\n"
	locGrp := rxMustCompile(`(\n[ \t]*)+`).FindAllStringIndex(xml, -1) // BLANK lines
	xml = replByPosGrp(xml, locGrp, []string{""})                      // remove all BLANK lines
	xml = xmlfmt.FormatXML(xml, "", "\t")                              // NOTICE: after this, auto "\r\n" applied by 'xmlfmt'
	xml = sReplaceAll(xml, "\r\n", "\n")                               // "\r\n" -> "\n"
	xml = sReplaceAll(xml, "    ", "\t")                               // "4space" -> "\t"

	indent, cnt := "", 0
	for i := 0; i < 100; i++ {
		r := rxMustCompile(fSf(`[^>]\n%s</`, indent))
		locGrp := r.FindAllStringIndex(xml, -1)
		if locGrp == nil {
			if cnt == 4 {
				break
			}
			cnt++
		} else {
			xml = replByPosGrp(xml, locGrp, []string{""}, 1, 2)
		}
		indent += "\t"
	}
	return sTrim(xml, " \t\r\n")
}

// XMLRoot :
func XMLRoot(xml string) string {
	xml = sTrim(xml, " \t\n\r")
	failP1OnErrWhen(!isXML(xml), "%v", n3err.PARAM_INVALID_XML)

	// normal tags style, "<tag (attr='abc')>value</tag>"
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
	if end > start {
		return xml[start:end]
	}

	// empty value, short style, with attributes, "<tag attr='abc'/>"
	segs := sSplit(xml, " ")
	if len(segs) > 1 {
		return segs[0][1:]
	}

	// empty value, short style, without attributes, "<tag/>"
	return xml[1 : len(xml)-2]
}

// HasAllAttr :
func HasAllAttr(xml string, attrs ...string) bool {
	if (sCount(xml, "<") > 2 && sCount(xml, ">") > 2) || !isXML(xml) {
		return false
	}
	for _, attr := range attrs {
		if !rxMustCompile(fSf(` %s=["'].*["'][ >/]`, attr)).MatchString(xml) {
			return false
		}
	}
	return true
}

// HasAnyAttr :
func HasAnyAttr(xml string, attrs ...string) bool {
	if (sCount(xml, "<") > 2 && sCount(xml, ">") > 2) || !isXML(xml) {
		return false
	}
	for _, attr := range attrs {
		if rxMustCompile(fSf(` %s=["'].*["'][ >/]`, attr)).MatchString(xml) {
			return true
		}
	}
	return false
}

// AttrValue :
func AttrValue(xml string) (attrs []string, mAttrVal map[string]string) {
	xml = sTrim(xml, " \t\n\r")
	if (sCount(xml, "<") > 2 && sCount(xml, ">") > 2) || !isXML(xml) {
		return nil, nil
	}
	root := XMLRoot(xml)
	if hasAnyPrefix(xml, fSf(`<%s>`, root), fSf(`<%s/>`, root)) {
		return []string{}, map[string]string{}
	}

	mAttrVal = make(map[string]string)
	head, tail1, tail2 := fSf(`<%s `, root), fSf(`</%s>`, root), fSf(`/>`)
	attrseg := ""
	if loc := rxMustCompile(fSf(`>.*%s$`, tail1)).FindStringIndex(xml); loc != nil {
		attrseg = xml[len(head):loc[0]]
	}
	if sHasSuffix(xml, tail2) {
		attrseg = xml[len(head) : len(xml)-len(tail2)]
	}

	for _, avseg := range sSplit(attrseg, " ") {
		av := sSplit(avseg, "=")
		attrs = append(attrs, av[0])
		mAttrVal[av[0]] = av[1][1 : len(av[1])-1]
	}
	return
}

// Value :
func Value(xml string) (string, bool) {
	xml = sTrim(xml, " \t\n\r")
	if (sCount(xml, "<") > 2 && sCount(xml, ">") > 2) || !isXML(xml) {
		return "", false
	}
	if sHasSuffix(xml, "/>") {
		return "", true
	}
	if loc := rxMustCompile(`>.*<`).FindStringIndex(xml); loc != nil {
		return xml[loc[0]+1 : loc[1]-1], true
	}
	return "", true
}
