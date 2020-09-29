package n3xml

const (
	linker = "~"
	preatt = "-"
)

var mPathIdx = make(map[string]int)
var mIPathVal = make(map[string]string)

// Explicitly indicate which PATH is LIST
var mList = map[string]bool{
	"sif~Activity": true,
	"sif~Identity": true,
}

// var mIPathVal = sync.Map{}

// decode :
func decode(pPath, xml string) {
	if tag, cont, _, av := TagContAttrVal(xml); tag != "" {
		path := tag
		if pPath != "" {
			path = fSf("%s~%s", pPath, tag)
		}

		if _, ok := mPathIdx[path]; !ok {
			mPathIdx[path] = -1
		}
		mPathIdx[path]++

		idx := mPathIdx[path]
		ipath := path
		if idx > 0 {
			ipath = fSf("%s#%d", path, idx)
		} else {
			// enforce LIST
			if _, ok := mList[path]; ok {
				ipath = fSf("%s#0", path)
			}
		}

		// attributes :
		for k, v := range av {
			mIPathVal[ipath+"~-"+k] = v // "-" is marked as attribute
			// fPln(ipath+"~-"+k, v)
		}

		// content :
		mIPathVal[ipath] = cont
		// fPln(ipath, cont)

		// Next :
		_, subs := BreakCont(cont)
		for _, sub := range subs {
			decode(ipath, sub)
		}
	}

	return
}

// Decode :
func Decode(xml string) {
	decode("", xml)
}
