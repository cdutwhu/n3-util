package n3xml

import (
	"github.com/cdutwhu/n3-util/n3err"
	"github.com/go-xmlfmt/xmlfmt"
)

var (
	// Fmt
	rxHead        = rxMustCompile(`<\w+(\s+\w+\s*=\s*"[^"]*"\s*)*\s*/?>`)
	rxTail        = rxMustCompile(`</\w+\s*>`)
	rxContMultiLF = rxMustCompile(`(\n\t*){2,}</?\w+`)
	rxMultiLF     = rxMustCompile(`(\n\t*){2,}`)
	rxSTagLoose   = rxMustCompile(`<\w+\s{2,}(>|\w+)`)
	rxSTagSpace   = rxMustCompile(`\s{2,}`)
	rxAttrLoose   = rxMustCompile(`"\s{2,}[\w/>]`)
	rxAttrSpace   = rxMustCompile(`\s{2,}`)
	rxEqLoose     = rxMustCompile(`\s[^"=]+\s*=\s*"`)
	rxEqSpace     = rxMustCompile(`\s+=\s+`)
	rxAttrDangle  = rxMustCompile(`<\w+(\n\s*)+[^>]*>`)
	rxAttrLF      = rxMustCompile(`(\n\s*)+`)
	rxETagLoose   = rxMustCompile(`</\w+\s+>`)
	rxETagSpace   = rxMustCompile(`\s+`)
	// TagContAttrVal
	rxAttrPart = rxMustCompile(`\s+\w+\s*=\s*"[\w\s]*"`)
	rxSTag     = rxMustCompile(`<\w+[\s>/]`)
)

// Fmt :
func Fmt(xml string) string {
	xml = xmlfmt.FormatXML(xml, "", "\t") // NOTICE: after this, auto "\r\n" applied by 'xmlfmt'
	xml = sReplaceAll(xml, "\r\n", "\n")  // "\r\n" -> "\n"

	// return xml

	// Remove LF at end of content
	xml = rxContMultiLF.ReplaceAllStringFunc(xml, func(m string) string {
		s, e := -1, -1
		for i := len(m) - 1; i >= 0; i-- {
			if m[i] == '\t' {
				if e == -1 {
					e = i + 1
				}
				continue
			}
			if m[i] == '\n' {
				s = i
				break
			}
		}
		return rxMultiLF.ReplaceAllString(m, m[s:e])
	})

	mOldNew := make(map[string]string)
	search, I := xml, 0
NEXT1:
	if pair := rxHead.FindAllStringIndex(search, 1); pair != nil {
		s, e := pair[0][0], pair[0][1]
		find := search[s:e]
		// fPln(I, find)
		search = search[e:]

		rmAttrDangle := rxAttrDangle.ReplaceAllStringFunc(find, func(m string) string {
			return rxAttrLF.ReplaceAllString(m, " ")
		})
		//-----------------------------------//
		sTagFmt := rxSTagLoose.ReplaceAllStringFunc(rmAttrDangle, func(m string) string {
			return sReplaceAll(rxSTagSpace.ReplaceAllString(m, " "), " >", ">")
		})
		//-----------------------------------//
		attrFmt := rxAttrLoose.ReplaceAllStringFunc(sTagFmt, func(m string) string {
			return sReplaceAll(rxAttrSpace.ReplaceAllString(m, " "), "\" >", "\">")
		})
		//-----------------------------------//
		eqFmt := rxEqLoose.ReplaceAllStringFunc(attrFmt, func(m string) string {
			return rxEqSpace.ReplaceAllString(m, "=")
		})

		if eqFmt != find {
			mOldNew[find] = eqFmt
		}

		I++
		goto NEXT1
	}

	search, I = xml, 0
NEXT2:
	if pair := rxTail.FindAllStringIndex(search, 1); pair != nil {
		s, e := pair[0][0], pair[0][1]
		find := search[s:e]
		// fPln(I, find)
		search = search[e:]

		eTagFmt := rxETagLoose.ReplaceAllStringFunc(find, func(m string) string {
			return rxETagSpace.ReplaceAllString(m, "")
		})

		if eTagFmt != find {
			mOldNew[find] = eTagFmt
		}

		I++
		goto NEXT2
	}

	for k, v := range mOldNew {
		xml = sReplaceAll(xml, k, v)
	}

	return sTrim(xml, " \t\r\n")
}

// TagContAttrVal :
func TagContAttrVal(xml string) (tag, cont string, attrs []string, mAttrVal map[string]string) {

	xml = sTrim(xml, " \t\n\r")
	if !isXML(xml) {
		return "", "", nil, nil
	}

	sTag, eTag := "", ""
	if pair := rxHead.FindAllStringIndex(xml, 1); pair != nil {
		s, e := pair[0][0], pair[0][1]
		sTag = xml[s:e]
		// fPln(1, sTag)

		tag = rxSTag.FindAllString(sTag, 1)[0]
		tag = sTrim(tag, "< \t\r\n>/")
		// fPln(2, tag)

		mAttrVal = make(map[string]string)
		for _, av := range rxAttrPart.FindAllString(sTag, -1) {
			av = sTrim(av, " \t\r\n")
			ss := sSplit(av, "=")
			a, v := sTrim(ss[0], " \t\r\n"), sTrim(sTrim(ss[1], " \t\r\n"), "\"")
			attrs = append(attrs, a)
			mAttrVal[a] = v
		}
	}

	found := false
	if pairs := rxTail.FindAllStringIndex(xml, -1); pairs != nil {
		eTag1st := fSf("</%s>", tag)
		for _, pair := range pairs {
			s, e := pair[0], pair[1]
			eTag = xml[s:e]
			if eTag == eTag1st {
				found = true
				break
			}
		}
		// fPln(3, eTag)
	}

	if found {
		start := sIndex(xml, sTag) + len(sTag)
		cont = xml[start:]
		cont = cont[:sIndex(cont, eTag)]
		cont = sTrimRight(cont, " \t\r\n")
		// fPln(4, cont)
	}

	return
}

// XMLRoot :
func XMLRoot(xml string) string {
	xml = sTrim(xml, " \t\n\r")
	failP1OnErrWhen(!isXML(xml), "%v", n3err.PARAM_INVALID_XML)
	tag, _, _, _ := TagContAttrVal(xml)
	return tag
}
