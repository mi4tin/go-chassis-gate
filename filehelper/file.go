package filehelper

import (
	"fmt"
	"github.com/go-chassis/go-chassis/pkg/util/fileutil"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"path/filepath"
)

const (
	//FileNameGate gate配置文件名
	FileNameGate = "gate.yaml"
)

//GetConfig 获取配置，configObj必须为指定类型的struct
func GetConfig(configObj interface{}, fileName string) error {
	filePath := filepath.Join(fileutil.GetConfDir(), fileName)
	data, err := ioutil.ReadFile(filePath)
	fmt.Println("Config.err:", err)
	err = yaml.Unmarshal(data, configObj)
	return err
}
