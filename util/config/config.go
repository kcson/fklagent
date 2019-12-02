package config

import (
	"fmt"
	"github.com/magiconair/properties"
	"os"
	"path/filepath"
)

type config struct {
	DBDriver string `properties:"db.driver,default=postgres"`
	DBHost   string `properties:"db.host,default=211.251.244.64"`
	DBPort   int    `properties:"db.port,default=10038"`
	DBUser   string `properties:"db.user,default=bbp_prd_fasoo"`
	DBPass   string `properties:"db.pass,default=new1234!"`
	DBName   string `properties:"db.name,default=bbpa1"`

	SFTPPath string `properties:"data.sftp.path,default=/data1/dev/sftp"`
	ESBPath  string `properties:"data.esb.path,default=/data1/dev/esb"`

	RCmd        string `properties:"r.cmd"`
	RCmdParam   string `properties:"r.cmd.param"`
	RScriptPath string `properties:"r.script.path"`
	RResultPath string `properties:"r.result.path,default=/home/fasoo/R/result"`
	RLogPath    string `properties:"r.log.path,default=/home/fasoo/R/log"`
}

var Cfg config
var P *properties.Properties

func init() {
	fmt.Println("init config")
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	homePath := filepath.Dir(ex)
	fmt.Println(homePath)
	configPath := homePath + string(os.PathSeparator) + "config" + string(os.PathSeparator) + "fklagent.properties"

	P = properties.MustLoadFile(configPath, properties.UTF8)
	if err := P.Decode(&Cfg); err != nil {
		panic(err)
	}
}
