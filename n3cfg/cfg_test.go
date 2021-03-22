package n3cfg

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/davecgh/go-spew/spew"
)

func TestPrjName(t *testing.T) {
	SetDftCfgVal("dftPrjName", "0.0.0")
	fPln(PrjName())
}

func TestGitVer(t *testing.T) {
	SetDftCfgVal("dftPrjName", "0.0.0")
	fPln(GitVer())
}

func TestGitTag(t *testing.T) {
	fPln(GitTag())
}

func TestModify(t *testing.T) {
	cfg := &Config{}
	_, err := toml.DecodeFile("../data/toml/test.toml", cfg)
	failOnErr("%v", err)
	Icfg := Modify(cfg, map[string]interface{}{
		"[PORT]": 1234,
		"[s]":    "n3cfg",
		"[v]":    "v1.2.3",
	})
	cfg = Icfg.(*Config)
	spew.Dump(cfg)
}

func TestNewCfg(t *testing.T) {
	SetDftCfgVal("dftPrjName", "0.0.0")
	prj, _ := PrjName()
	cfg := &Config{}
	ok := New(
		cfg,
		map[string]string{
			"[s]":   "WebService.Service",
			"[v]":   "WebService.Version",
			"[p]":   "Port",
			"[prj]": prj,
		},
		"../data/toml/test.toml",
	)
	fPln(ok)
	Save("./saved.toml", cfg)
}

func TestEvalCfgValue(t *testing.T) {
	cfg := &Config{}
	ok := New(cfg, nil, "../data/toml/test.toml")
	fPln(ok)
	fPln(EvalCfgValue(cfg, "Port"))
	fPln(EvalCfgValue(cfg, "Storage.MetaPath"))
	fPln(EvalCfgValue(cfg, "WebService.Port"))
	fPln(EvalCfgValue(cfg, "Server.Service"))
	fPln(EvalCfgValue(cfg, "WebService.Service"))
	Save("./saved.toml", cfg)
}

func TestToEnvVar(t *testing.T) {
	// mReplExpr := map[string]string{
	// 	"[s]": "Service",
	// 	"[v]": "WebService.Version",
	// 	"[p]": "Port",
	// }
	// cfg := ToEnvN3utilServer(mReplExpr, "KEY", "../data/toml/test.toml")
	// spew.Dump(cfg)
	// fPln(" ----------------------------------------------------------- ")
	// cfg1 := FromEnvN3utilServer("KEY")
	// if cfg1 == nil {
	// 	fPln("Error @ FromEnvVar")
	// 	return
	// }
	// spew.Dump(cfg1)
}

// echo 'password' | sudo -S env "PATH=$PATH" go test -v -count=1 ./ -run TestRegister
func TestRegister(t *testing.T) {
	prj, _ := PrjName()
	pkg := "Server"
	ok, file := Register("qmiao", "../data/toml/test.toml", prj, pkg)
	fPln(ok, file)
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
// 		files, _ := os.ReadDir(path)
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
