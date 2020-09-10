package toml

import "github.com/cdutwhu/n3-util/n3cfg"

// Config1 : AUTO Created From /home/qmiao/Desktop/n3-util/data/toml/test.toml
type Config1 struct {
	IP      interface{}
	LogFile string
	Path    interface{}
	Port    int
	Service interface{}
	Storage struct {
		BadgerDBPath string
		DataBase     string
		MetaPath     string
	}
	WebService struct {
		Port    interface{}
		Service string
		Version interface{}
	}
	Route struct {
		Get       string
		GetHash   string
		GetID     string
		LsUser    string
		Delete    string
		LsContext string
		LsID      string
		LsObject  string
		ROOT      string
		Update    string
	}
	File struct {
		ClientMac     string
		ClientWin64   string
		MaskConfig    string
		MaskLinux64   string
		MaskMac       string
		MaskWin64     string
		ClientConfig  string
		ClientLinux64 string
	}
	Server struct {
		IP       interface{}
		Port     interface{}
		Protocol string
		Service  string
	}
	Access struct {
		Timeout int
	}
	SchoolPrograms struct {
		BOOLEAN []string
	}
	CollectionStatus struct {
		NUMERIC []string
	}
	Journal struct {
		NUMERIC []string
	}
	SectionInfo struct {
		NUMERIC []string
	}
}

// Config2 : AUTO Created From /home/qmiao/Desktop/n3-util/data/toml/test1.toml
type Config2 struct {
	Version string
	Storage struct {
		BadgerDBPath string
		DataBase     string
		MetaPath     string
	}
	WebService struct {
		Version interface{}
		Port    interface{}
		Service string
	}
	Route struct {
		Update    string
		GetHash   string
		GetID     string
		LsContext string
		LsObject  string
		LsUser    string
		ROOT      string
		Delete    string
		Get       string
		LsID      string
	}
	File struct {
		MaskConfig    string
		MaskLinux64   string
		MaskMac       string
		MaskWin64     string
		ClientConfig  string
		ClientLinux64 string
		ClientMac     string
		ClientWin64   string
	}
	Server struct {
		IP       interface{}
		Port     interface{}
		Protocol string
		Service  string
	}
	Access struct {
		Timeout int
	}
	SchoolPrograms struct {
		BOOLEAN []string
	}
	CollectionStatus struct {
		NUMERIC []string
	}
	Journal struct {
		NUMERIC []string
	}
	SectionInfo struct {
		NUMERIC []string
	}
}

// NewCfg :
func NewCfg(cfgStruName string, mReplExpr map[string]string, cfgPaths ...string) interface{} {
	var cfg interface{}
	switch cfgStruName {
	case "Config1":
		cfg = &Config1{}
	case "Config2":
		cfg = &Config2{}
	default:
		return nil
	}
	return n3cfg.InitEnvVar(cfg, mReplExpr, cfgStruName, cfgPaths...)
}
