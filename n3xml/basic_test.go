package n3xml

import (
	"io/ioutil"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	defer trackTime(time.Now())

	xml := `<text a =  "23" >Thank you for sending us the information on
	<emphasis>SDL Trados Studio 2015 SDL Trados Studio 2015
	</emphasis>. 
</text>`

	//  xml := `<Duration   Units  =   "    minute   "   Unit="123">3220
	//  3221   11
	//  </Duration>`

	// xml := `<Duration>30007</Duration>`
	// xml := `<Duration Units="minute" Unit="123"/>`
	// xml := `<Duration/>`
	// xml := `<ActivityTime>`

	fPln(XMLRoot(xml))

	tag, cont, attrs, mAV := TagContAttrVal(xml)
	fPln(tag)
	fPln(cont)
	fPln(attrs)
	fPln(mAV)
	// fPln(exist("Unit", toGeneralSlc(attrs)...))

	return
}

func TestDigitalTags(t *testing.T) {
	// defer trackTime(time.Now())
	// bytes, err := ioutil.ReadFile("../data/xml/siftest347.xml")
	// failOnErr("%v", err)
	// r := rxMustCompile(`[0-9]*\.[0-9]{1}$`)
	// for _, ln := range splitLn(string(bytes)) {
	// 	if val, ok := Value(ln); ok && isNumeric(val) {
	// 		if r.MatchString(val) {
	// 			fPln(ln)
	// 		}
	// 	}
	// }
}

func TestFmt(t *testing.T) {
	defer trackTime(time.Now())
	bytes, err := ioutil.ReadFile("../data/xml/d.xml")
	failOnErr("%v", err)
	fPln(isXML(string(bytes)))
	xml := Fmt(string(bytes))
	fPln(isXML(xml))
	mustWriteFile("../data/xml/fmt.xml", []byte(xml))
}
