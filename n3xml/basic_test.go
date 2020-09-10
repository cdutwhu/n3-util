package n3xml

import (
	"io/ioutil"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	defer trackTime(time.Now())

	xml := `<Duration Units="minute" Unit="123">3220</Duration>`
	fPln(HasAllAttr(xml, "Unit", "Units"))
	fPln(HasAnyAttr(xml, "Unit", "abc"))
	fPln(AttrValue(xml))
	fPln(Value(xml))
	fPln("--------------------")

	xml = `<Duration>30007</Duration>`
	fPln(HasAllAttr(xml, "Unit", "Units"))
	fPln(HasAnyAttr(xml, "Unit", "abc"))
	fPln(AttrValue(xml))
	fPln(Value(xml))
	fPln("--------------------")

	xml = `<Duration Units="minute" Unit="123"/>`
	fPln(HasAllAttr(xml, "Unit", "Units"))
	fPln(HasAnyAttr(xml, "Unit", "abc"))
	fPln(AttrValue(xml))
	fPln(Value(xml))
	fPln("--------------------")

	xml = `<Duration/>`
	fPln(HasAllAttr(xml, "Unit", "Units"))
	fPln(HasAnyAttr(xml, "Unit", "abc"))
	fPln(AttrValue(xml))
	fPln(Value(xml))
	fPln("--------------------")

	xml = `<ActivityTime>`
	fPln(HasAllAttr(xml, "Unit", "Units"))
	fPln(HasAnyAttr(xml, "Unit", "abc"))
	fPln(AttrValue(xml))
	fPln(Value(xml))
	fPln("--------------------")
}

func TestDigitalTags(t *testing.T) {
	defer trackTime(time.Now())
	bytes, err := ioutil.ReadFile("../data/xml/siftest347.xml")
	failOnErr("%v", err)
	r := rxMustCompile(`[0-9]*\.[0-9]{1}$`)
	for _, ln := range splitLn(string(bytes)) {
		if val, ok := Value(ln); ok && isNumeric(val) {
			if r.MatchString(val) {
				fPln(ln)
			}
		}
	}
}

func TestFmt(t *testing.T) {
	defer trackTime(time.Now())
	bytes, err := ioutil.ReadFile("../data/xml/c.xml")
	failOnErr("%v", err)
	xml := Fmt(string(bytes))	
	mustWriteFile("../data/xml/fmt.xml", []byte(xml))
}
