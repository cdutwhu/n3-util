package common

import (
	"testing"
	"time"

	"github.com/burntsushi/toml"
)

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

func TestCfgRepl(t *testing.T) {
	cfg := &Config{}
	toml.DecodeFile("../_data/toml/test.toml", cfg)
	Icfg, err := CfgRepl(cfg, map[string]interface{}{
		"[DATE]": time.Now().Format("2006-01-02"),
		"[IP]":   LocalIP(),
		"[PORT]": cfg.WebService.Port,
		"[s]":    cfg.WebService.Service,
		"[v]":    cfg.WebService.Version,
	})
	FailOnErr("%v", err)
	cfg = Icfg.(*Config)
	fPln(cfg.Server.Port.(int) + 111)
	fPln(*cfg)
}
