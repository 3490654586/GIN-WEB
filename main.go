package main

import (
	"wep-app/dao/mysql"
	"wep-app/logger"
	"wep-app/router"
	"wep-app/utli"
)




func main ()  {
	utli.Init("config.json")
	logger.Init("debug")
	mysql.Init()
	router.Init("debug")
}


