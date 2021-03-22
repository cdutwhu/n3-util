package n3xml

import (
	"os"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
)

func TestBreak(t *testing.T) {
	defer misc.TrackTime(time.Now())
	bytes, err := os.ReadFile("../data/xml/sif.xml")
	failOnErr("%v", err)

	// xml := `<root> Thank you for sending us the information on
	// <emphasis b= "ss0" > SDL0 Trados Studio 2015 SDL Trados Studio 2015
	// <emphasis1 b  = "ss-111222"  />
	// </emphasis> ?????
	// <emphasis b= "ss-1"/>
	// <emphasis b= "ss1" > SDL1 Trados Studio 2015 SDL Trados Studio 2016
	// <test/>
	// </emphasis>. Hello
	// <test  /></root>`

	_, ok := Break(string(bytes), "./out", true)
	fPln(ok)
}

func TestSmashCont(t *testing.T) {
	defer misc.TrackTime(time.Now())

	// xml := ` Thank you for sending us the information on
	// <emphasis b= "ss0" > SDL0 Trados Studio 2015 SDL Trados Studio 2015
	// <emphasis1 b= "ss-111"/>
	// </emphasis> ?????
	// <emphasis b= "ss-1"/>
	// <emphasis b= "ss1" > SDL1 Trados Studio 2015 SDL Trados Studio 2016
	// <test/>
	// </emphasis>. Hello
	// <test  />`

	// bytes, err := os.ReadFile("./sif/NAPCodeFrame_0.xml")
	bytes, err := os.ReadFile("../data/xml/sif.xml")
	failOnErr("%v", err)
	xml := string(bytes)

	tag, cont, attrs, m := TagContAttrVal(xml)
	fPln(tag)
	fPln(attrs)
	fPln(m)

	fPln(" ----------------------------------- ")

	roots, subs := BreakCont(cont)
	fPln(len(roots))
	// for _, root := range roots {
	// 	fPln(root)
	// }
	// for _, sub := range subs {
	// 	fPln(sub)
	// 	fPln(" --- ")
	// }
	fPln(subs[0])
}

func TestFastSmashCont(t *testing.T) {
	defer misc.TrackTime(time.Now())

	bytes, err := os.ReadFile("../data/xml/fmt.xml")
	failOnErr("%v", err)
	xml := string(bytes)
	tag, cont, attrs, m := TagContAttrVal(xml)
	fPln(tag)
	fPln(attrs)
	fPln(m)

	fPln(" ----------------------------------- ")

	roots, subs := FastBreakCont(cont)
	fPln(len(roots))
	// for _, root := range roots {
	// 	fPln(root)
	// }
	fPln(subs[0])
}
