package n3cfg

import (
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

// Config : holder for ignoring compiling error. genStruct to get an real one.
// then comment below out.
// type Config struct {
// }

func TestModify(t *testing.T) {
	cfg := &Config{}
	_, err := toml.DecodeFile("../_data/toml/test.toml", cfg)
	failOnErr("%v", err)
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

func TestNewCfg(t *testing.T) {
	cfg := &Config{}
	ok := New(
		cfg,
		map[string]string{
			"[s]": "WebService.Service",
			"[v]": "WebService.Version",
			"[p]": "Port",
		},
		"../_data/toml/test.toml",
	)
	fPln(ok)
	spew.Dump(cfg)
	Save("./saved.toml", cfg)
}

func TestEvalCfgValue(t *testing.T) {
	cfg := &Config{}
	ok := New(cfg, nil, "../_data/toml/test.toml")
	fPln(ok)
	fPln(EvalCfgValue(cfg, "Port"))
	fPln(EvalCfgValue(cfg, "Storage.MetaPath"))
	fPln(EvalCfgValue(cfg, "WebService.Port"))
	fPln(EvalCfgValue(cfg, "Server.Service"))
	fPln(EvalCfgValue(cfg, "WebService.Service"))
	Save("./saved.toml", cfg)
}

// func TestWorldTime(t *testing.T) {
// 	tmin := func(t time.Time, name string) (time.Time, error) {
// 		loc, err := time.LoadLocation(name)
// 		if err == nil {
// 			t = t.In(loc)
// 		}
// 		return t, err
// 	}

// 	for _, name := range []string{
// 		"",
// 		"Local",
// 		"Asia/Shanghai",
// 		"America/New_York",
// 		"Australia/Melbourne",
// 	} {
// 		t, err := tmin(time.Now(), name)
// 		if err == nil {
// 			fPln(t.Location(), t.Format("15:04"))
// 		} else {
// 			fPln(name, "<time unknown>")
// 		}
// 	}

// 	fPln(" --------------------------------------- ")

// 	now := time.Now()
// 	zone, offset := now.Zone()
// 	fPln(zone, offset)

// 	fPln(" --------------------------------------- ")
// }

// func TestListAllLoc(t *testing.T) {

// 	var readFile func(string)

// 	readFile = func(path string) {
// 		files, _ := ioutil.ReadDir(path)
// 		for _, f := range files {
// 			if f.Name() != sToUpper(f.Name()[:1])+f.Name()[1:] {
// 				continue
// 			}
// 			if f.IsDir() {
// 				readFile(path + "/" + f.Name())
// 			} else {
// 				fPln((path + "/" + f.Name())[1:])
// 			}
// 		}
// 	}

// 	for _, zoneDir := range []string{
// 		// Update path according to your OS
// 		"/usr/share/zoneinfo/",
// 		"/usr/share/lib/zoneinfo/",
// 		"/usr/lib/locale/TZ/",
// 	} {
// 		readFile(zoneDir)
// 	}
// }
