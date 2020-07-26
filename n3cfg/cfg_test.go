package n3cfg

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/davecgh/go-spew/spew"
)

func TestGitVer(t *testing.T) {
	fPln(GitVer())
}

func TestGitTag(t *testing.T) {
	fPln(GitTag())
}

// 'genStruct' to create one for real one
type Config struct {
}

func TestModify(t *testing.T) {
	cfg := &Config{}
	toml.DecodeFile("../_data/toml/test.toml", cfg)
	Icfg := Modify(cfg, map[string]interface{}{
		"[DATE]": time.Now().Format("2006-01-02"),
		"[IP]":   localIP(),
		"[PORT]": 1234,
		"[s]":    "n3cfg",
		"[v]":    "v1.2.3",
	})
	cfg = Icfg.(*Config)
	spew.Dump(cfg)
}

func TestWorldTime(t *testing.T) {
	tmin := func(t time.Time, name string) (time.Time, error) {
		loc, err := time.LoadLocation(name)
		if err == nil {
			t = t.In(loc)
		}
		return t, err
	}

	for _, name := range []string{
		"",
		"Local",
		"Asia/Shanghai",
		"America/New_York",
		"Australia/Melbourne",
	} {
		t, err := tmin(time.Now(), name)
		if err == nil {
			fPln(t.Location(), t.Format("15:04"))
		} else {
			fPln(name, "<time unknown>")
		}
	}

	fPln(" --------------------------------------- ")

	now := time.Now()
	zone, offset := now.Zone()
	fPln(zone, offset)

	fPln(" --------------------------------------- ")
}

func TestListAllLoc(t *testing.T) {

	var readFile func(string)

	readFile = func(path string) {
		files, _ := ioutil.ReadDir(path)
		for _, f := range files {
			if f.Name() != sToUpper(f.Name()[:1])+f.Name()[1:] {
				continue
			}
			if f.IsDir() {
				readFile(path + "/" + f.Name())
			} else {
				fPln((path + "/" + f.Name())[1:])
			}
		}
	}

	for _, zoneDir := range []string{
		// Update path according to your OS
		"/usr/share/zoneinfo/",
		"/usr/share/lib/zoneinfo/",
		"/usr/lib/locale/TZ/",
	} {
		readFile(zoneDir)
	}
}