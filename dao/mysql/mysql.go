package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"wep-app/model"
	"wep-app/utli"
)


var Db *xorm.Engine
var err error
func Init()  {
	//_conf := utli.Init("../../config.json")
	_conf := utli.GetConf()
	fmt.Println(1)
	fmt.Println("_conf=",_conf.Mysqls.Databases)
	Db,err = xorm.NewEngine(_conf.Mysqls.Databases,fmt.Sprint(_conf.Mysqls.User+":"+_conf.Mysqls.Passwad+"@tcp("+_conf.Mysqls.Host+":"+_conf.Mysqls.Port+")/"+_conf.Mysqls.Db_Name))
	Db.Sync2(new(model.User))
}
