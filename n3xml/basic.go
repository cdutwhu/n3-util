package n3xml

import (
	"github.com/cdutwhu/n3-util/n3err"
	"github.com/go-xmlfmt/xmlfmt"
)

var (
	rxHead    = rxMustCompile(`<\w+(\s+\w+\s*=\s*"[^"]*"\s*)*\s*/?>`)
	rxTail    = rxMustCompile(`</\w+\s*>`)
	rxMultiLF = rxMustCompile(`\n{2}\t+<\w+`)
	// rxLF         = rxMustCompile(`\n{2}\t+`)
	rxTagLoose   = rxMustCompile(`<\w+\s{2,}(>|\w+)`)
	rxTagSpace   = rxMustCompile(`\s{2,}`)
	rxAttrLoose  = rxMustCompile(`"\s{2,}[\w/>]`)
	rxAttrSpace  = rxMustCompile(`\s{2,}`)
	rxEqLoose    = rxMustCompile(`\s[^"=]+\s*=\s*"`)
	rxEqSpace    = rxMustCompile(`\s+=\s+`)
	rxAttrDangle = rxMustCompile(`<\w+(\n\s*)+[^>]*>`)
	rxAttrLF     = rxMustCompile(`(\n\s*)+`)
)

// Fmt :
func Fmt(xml string) string {
	xml = xmlfmt.FormatXML(xml, "", "\t") // NOTICE: after this, auto "\r\n" applied by 'xmlfmt'
	xml = sReplaceAll(xml, "\r\n", "\n")  // "\r\n" -> "\n"

	// Remove new added LF when tail is already LF
	xml = rxMultiLF.ReplaceAllStringFunc(xml, func(m string) string {
		return sReplace(m, "\n\n", "\n", 1)
	})

	mOldNew := make(map[string]string)
	search := xml
	I := 0
NEXT:
	if pair := rxHead.FindAllStringIndex(search, 1); pair != nil {
		s, e := pair[0][0], pair[0][1]
		find := search[s:e]
		fPln(I, find)
		search = search[e:]

		rmAttrDangle := rxAttrDangle.ReplaceAllStringFunc(find, func(m string) string {
			return rxAttrLF.ReplaceAllString(m, " ")
		})
		//-----------------------------------//
		tagThin := rxTagLoose.ReplaceAllStringFunc(rmAttrDangle, func(m string) string {
			return sReplaceAll(rxTagSpace.ReplaceAllString(m, " "), " >", ">")
		})
		//-----------------------------------//
		attrThin := rxAttrLoose.ReplaceAllStringFunc(tagThin, func(m string) string {
			return sReplaceAll(rxAttrSpace.ReplaceAllString(m, " "), "\" >", "\">")
		})
		//-----------------------------------//
		eqThin := rxEqLoose.ReplaceAllStringFunc(attrThin, func(m string) string {
			return rxEqSpace.ReplaceAllString(m, "=")
		})

		if eqThin != find {
			mOldNew[find] = eqThin
			fPln(I, find)
			fPln(I, eqThin)
			fPln(" ----------------------- ")
		}

		I++
		goto NEXT
	}

	for k, v := range mOldNew {
		xml = sReplaceAll(xml, k, v)
	}

	// if pair := rxTail

	// fPln(" ----------------------- ")
	// for i, tail := range rxTail.FindAllString(xml, -1) {
	// 	fPln(i, tail)
	// }

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
	if loc := rxMustCompile(fSf(`>([^<>]*(\n\s*)*)+%s$`, tail1)).FindStringIndex(xml); loc != nil {
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
	if loc := rxMustCompile(`>([^<>]*(\n\s*)*)+<`).FindStringIndex(xml); loc != nil {
		return sTrim(xml[loc[0]+1:loc[1]-1], " \t\n\r"), true
	}
	return "", true
}
