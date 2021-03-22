package n3xml

import (
	"os"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
)

func TestDecode(t *testing.T) {
	defer misc.TrackTime(time.Now())
	bytes, err := os.ReadFile("../data/xml/sif.xml")
	failOnErr("%v", err)

	Decode(string(bytes))

	fPln(mIPathVal["NAPResultsReporting~SchoolInfo~SchoolName"])
}
