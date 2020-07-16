package cfg

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

// Config is toml
type Config struct {
	Path    string
	LogFile string

	Storage struct {
		DataBase     string
		MetaPath     string
		BadgerDBPath string
	}

	WebService struct {
		Port    int
		Service string
		Version string
	}

	Route struct {
		HELP        string
		GetID       string
		GetHash     string
		Get         string
		Update      string
		Delete      string
		LsID        string
		LsUser      string
		LsContext   string
		LsObject    string
		GetEnforced string
	}

	File struct {
		ClientLinux64   string
		ClientMac       string
		ClientWin64     string
		ClientConfig    string
		EnforcerLinux64 string
		EnforcerMac     string
		EnforcerWin64   string
		EnforcerConfig  string
	}

	Server struct {
		Protocol string
		IP       string
		Port     interface{}
	}

	Access struct {
		Timeout int
	}
}

func TestModify(t *testing.T) {
	cfg := &Config{}
	toml.DecodeFile("../_data/toml/test.toml", cfg)
	Icfg := Modify(cfg, map[string]interface{}{
		"[DATE]": time.Now().Format("2006-01-02"),
		"[IP]":   localIP(),
		"[PORT]": cfg.WebService.Port,
		"[s]":    cfg.WebService.Service,
		"[v]":    cfg.WebService.Version,
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
