package gate

import "github.com/mi4tin/go-chassis-gate/filehelper"

var configObj *Config

//Config 是gate相关的一些配置
type Config struct {
	//白名单
	IPWhiteList string `yaml:"ipWhiteList"`
}

func init() {
	initConfig()
}

//配置初始化
func initConfig() {
	configObj = &Config{}
	err := filehelper.GetConfig(configObj, filehelper.FileNameGate)
	if err != nil {
		panic(err)
	}
}

//GetConfig 获取配置
func GetConfig() *Config {
	return configObj
}
