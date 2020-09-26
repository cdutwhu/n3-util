package n3xml

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
)

func TestSmashFirstAndSave(t *testing.T) {
	defer misc.TrackTime(time.Now())
	bytes, err := ioutil.ReadFile("../data/xml/sif.xml")
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

	_, ok := Break(string(bytes), "./sif", false)
	fPln(ok)
}

func TestSmashCont(t *testing.T) {
	defer misc.TrackTime(time.Now())

	xml := ` Thank you for sending us the information on
	<emphasis b= "ss0" > SDL0 Trados Studio 2015 SDL Trados Studio 2015
	<emphasis1 b= "ss-111"/>
	</emphasis> ?????
	<emphasis b= "ss-1"/>
	<emphasis b= "ss1" > SDL1 Trados Studio 2015 SDL Trados Studio 2016
	<test/>
	</emphasis>. Hello 
	<test  />`

	bytes, err := ioutil.ReadFile("./sif/NAPCodeFrame_0.xml")
	failOnErr("%v", err)
	xml = string(bytes)
	tag, cont, attrs, m := TagContAttrVal(xml)
	fPln(tag)
	// fPln(cont)
	fPln(attrs)
	fPln(m)

	fPln(" ----------------------------------- ")

	roots, subs := SmashCont(cont)
	fPln(roots)
	fPln(" ----------------------------------- ")
	for _, sub := range subs {
		fPln(sub)
		fPln(" ------------- ")
	}
}
