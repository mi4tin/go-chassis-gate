package restful

import "github.com/go-chassis/go-chassis/gate/fileHelper"

var configObj *Config

type Config struct {
	IpWhiteList string  `yaml:"ipWhiteList"`
}

func init(){
	initConfig()
}

func initConfig() {
	configObj=&Config{}
	err:=fileHelper.GetConfig(configObj,fileHelper.FileName_Gate)
	if err!=nil{
		panic(err)
	}
}

func GetConfig() *Config{
	return configObj
}

