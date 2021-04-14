package server

import (
	"fmt"
	"testing"
	"wep-app/dao/mysql"
	"wep-app/model"
)

func TestCheckUser(t *testing.T) {
   mysql.Init()
   user :=model.User{
	   UserName:   "dadw",
	   UserPaword:"dwada",
   }
   fmt.Println(CheckUser(user))
}