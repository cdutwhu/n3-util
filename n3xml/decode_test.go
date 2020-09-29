package n3xml

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
)

func TestDecode(t *testing.T) {
	defer misc.TrackTime(time.Now())
	bytes, err := ioutil.ReadFile("../data/xml/sif.xml")
	failOnErr("%v", err)

	Decode(string(bytes))

	fPln(mIPathVal["NAPResultsReporting~SchoolInfo~SchoolName"])
}
