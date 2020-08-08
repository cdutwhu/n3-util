package n3cfg

// Config : AUTO Created From /home/qmiao/Desktop/n3-util/_data/toml/test.toml
type Config struct {
    Path interface{}
    Port int
    Service interface{}
    IP interface{}
    LogFile string
    Storage struct {
        MetaPath string
        BadgerDBPath string
        DataBase string
    }
    WebService struct {
        Version interface{}
        Port interface{}
        Service string
    }
    Route struct {
        GetID string
        LsContext string
        ROOT string
        Update string
        GetHash string
        Get string
        LsID string
        LsObject string
        LsUser string
        Delete string
    }
    File struct {
        MaskConfig string
        MaskLinux64 string
        MaskMac string
        MaskWin64 string
        ClientConfig string
        ClientLinux64 string
        ClientMac string
        ClientWin64 string
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
}
