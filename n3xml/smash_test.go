package n3xml

import (
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
)

func TestSmashFirstAndSave(t *testing.T) {
	// bytes, err := ioutil.ReadFile("../data/xml/siftest346.xml")
	// failOnErr("%v", err)

	xml := `<root> Thank you for sending us the information on
	<emphasis b= "ss0" > SDL0 Trados Studio 2015 SDL Trados Studio 2015
	<emphasis1 b  = "ss-111222"  />
	</emphasis> ?????
	<emphasis b= "ss-1"/>
	<emphasis b= "ss1" > SDL1 Trados Studio 2015 SDL Trados Studio 2016
	<test/>
	</emphasis>. Hello 
	<test  /></root>`

	fPln(SmashFirstAndSave(xml, "./sif346"))
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
	roots, subs := SmashCont(xml)
	fPln(roots)
	fPln(" ----------------------------------- ")
	for _, sub := range subs {
		fPln(sub)
		fPln(" ------------- ")
	}
}
