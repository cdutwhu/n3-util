package n3xml

import (
	"io/ioutil"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	defer trackTime(time.Now())

	xml := `abc <text a =  "23" > Thank you for sending us the information on
		<emphasis b= "ss" >SDL Trados Studio 2015 SDL Trados Studio 2015
		</emphasis>.
		<emphasis1 b= "ss" >SDL Trados Studio 2015 SDL Trados Studio 2016
		</emphasis1>.
	</text>
	<Duration   Units  =   "    minute   "   Unit="123">3220
	3221   11</Duration> def`

	// xml := `<Duration   Units  =   "    minute   "   Unit="123">3220
	//  3221   11</Duration>`

	// xml := `<Duration>30007</Duration>`
	// xml := `<Duration Units="minute" Unit="123"/>`
	// xml := `<Duration a =  "23" /><ActivityTime>abc</ActivityTime>`
	// xml := `<ActivityTime>`

	// fPln(XMLRoot(xml))

	tag, cont, attrs, mAV := TagContAttrVal(xml)
	fPln(tag)
	fPln(cont)
	fPln(attrs)
	fPln(mAV)
	// fPln(exist("Unit", toGeneralSlc(attrs)...))

	fPln(" --------------------------- ")

	tag, cont, attrs, mAV = TagContAttrVal(cont)
	fPln(tag)
	fPln(cont)
	fPln(attrs)
	fPln(mAV)

	fPln(" --------------------------- ")

	tag, cont, attrs, mAV = TagContAttrVal(cont)
	fPln(tag)
	fPln(cont)
	fPln(attrs)
	fPln(mAV)

	return
}

func TestDigitalTags(t *testing.T) {
	defer trackTime(time.Now())
	bytes, err := ioutil.ReadFile("../data/xml/siftest347.xml")
	failOnErr("%v", err)
	r := rxMustCompile(`\d*\.\d{2}$`)
	for _, ln := range splitLn(string(bytes)) {
		if _, val, _, _ := TagContAttrVal(ln); isNumeric(val) {
			if r.MatchString(val) {
				fPln(ln)
			}
		}
	}
}

func TestFmt(t *testing.T) {
	defer trackTime(time.Now())
	bytes, err := ioutil.ReadFile("../data/xml/sif.xml")
	failOnErr("%v", err)
	fPln(isXML(string(bytes)))
	xml := Fmt(string(bytes))
	fPln(isXML(xml))
	mustWriteFile("../data/xml/fmt.xml", []byte(xml))
}
