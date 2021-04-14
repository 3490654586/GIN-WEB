package utli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var _cof *Conf
/**
 *项目配置
 */
type Conf struct {
	App      string `json:"app"`
	AppModel string `json:"app_model"`
	AppHost  string `json:"app_host"`
	AppPost  string `json:"app_post"`
	Mysqls   Mysql  `json:"mysql"`
	Logs    Log `json:"log"`
}


/**
 *Mysql配置
 */
type Mysql struct {
	Databases string `json:"databases"`
	User      string `json:"user"`
	Passwad   string  `json:"passwad"`
	Host      string `json:"host"`
	Port      string `json:"port"`
	Db_Name   string `json:"db_name"`
}


//日志配置
type Log struct {
	Level string `json:"level"`
	Filename string `json:"filename"`
	MaxSize int `json:"max_size"`
	MaxAge int `json:"max_age"`
	MaxBackups int `json:"max_backups"`
}

func Init(path string) *Conf{
	file,err := os.Open(path)
   fmt.Println("文件",file)
	reader := bufio.NewReader(file)

	decoder := json.NewDecoder(reader)

	err = decoder.Decode(&_cof)

	if err!=nil {
		return nil
	}
	return _cof
}

func GetConf () *Conf {
	return _cof
}

