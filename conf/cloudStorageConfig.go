package conf

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

const CONF_NAME = "./cloudStorageConf.xml"

var (
	CSConf = new(CloudStorageConf)
)

type CloudStorageConf struct {
	Ip    string            `xml:"ip"`   // service listen on
	Port  int               `xml:"port"` // service listen on
	Etcd  []string          `xml:"etcd"` // etcd ip:port
	Log   cloudStorageLog   `xml:"log"`
	Mysql cloudStorageMysql `xml:"mysql"`
}

type cloudStorageMysql struct {
	Host     string `xml:"host"`
	Port     int    `xml:"port"`
	User     string `xml:"user"`
	Password string `xml:"password"`
	DBName   string `xml:"dbName"`
}

type cloudStorageLog struct {
	LogName    string `xml:"logName"`
	Name       string `xml:"name"`       // log name of project
	MaxSize    int    `xml:"maxSize"`    // max size of log (MB)
	MaxBackups int    `xml:"maxBackups"` // max number of old log
	MaxAge     int    `xml:"maxAge"`     // ax days of old log retained
	Compress   bool   `xml:"compress"`   // compress or not
	Level      string `xml:"level"`      // log level
}

func (conf *CloudStorageConf) newConf() (err error) {
	var (
		confBates []byte
	)
	if confBates, err = ioutil.ReadFile(CONF_NAME); err != nil {
		fmt.Println("read config file error:", err)
		return
	}

	if err = xml.Unmarshal(confBates, conf); err != nil {
		fmt.Println("conf unmarshal to struct error:", err)
		return
	}

	return nil
}

func CloudStorageConfInit() (err error) {
	if err = CSConf.newConf(); err != nil {
		return
	}

	return nil
}
