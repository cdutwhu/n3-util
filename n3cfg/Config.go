package n3cfg

// Config : AUTO Created From /home/qmiao/Desktop/temp/n3-util/data/toml/test.toml
type Config struct {
	IP interface{}
	LogFile string
	Path interface{}
	Port int
	Service interface{}
	Storage struct {
		BadgerDBPath string
		DataBase string
		MetaPath string
	}
	WebService struct {
		Port interface{}
		Service string
		Version interface{}
	}
	Route struct {
		LsObject string
		LsUser string
		ROOT string
		GetID string
		LsID string
		GetHash string
		LsContext string
		Update string
		Delete string
		Get string
	}
	File struct {
		ClientConfig string
		ClientLinux64 string
		ClientMac string
		ClientWin64 string
		MaskConfig string
		MaskLinux64 string
		MaskMac string
		MaskWin64 string
	}
	Server struct {
		Protocol string
		Service string
		IP interface{}
		Port interface{}
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

// Config1 : AUTO Created From /home/qmiao/Desktop/temp/n3-util/data/toml/test1.toml
type Config1 struct {
	Version string
	Storage struct {
		BadgerDBPath string
		DataBase string
		MetaPath string
	}
	WebService struct {
		Service string
		Version interface{}
		Port interface{}
	}
	Route struct {
		Delete string
		Get string
		GetID string
		LsID string
		LsObject string
		LsUser string
		GetHash string
		LsContext string
		ROOT string
		Update string
	}
	File struct {
		ClientWin64 string
		MaskConfig string
		MaskLinux64 string
		MaskMac string
		MaskWin64 string
		ClientConfig string
		ClientLinux64 string
		ClientMac string
	}
	Server struct {
		IP interface{}
		Port interface{}
		Protocol string
		Service string
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
	case "Config":
		cfg = &Config{}
	case "Config1":
		cfg = &Config1{}
	default:
		return nil
	}
	return InitEnvVar(cfg, mReplExpr, cfgStruName, cfgPaths...)
}
