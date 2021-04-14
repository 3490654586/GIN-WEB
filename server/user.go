package server

import (
	"fmt"
	"wep-app/dao/mysql"
	"wep-app/model"
	"wep-app/utli/errmsg"
)




//用户登陆
func CheckUser(user model.User)string{
	 var userModel model.User
      mysql.Db.Where("user_name = ? ",user.UserName).Get(&userModel)
	  if userModel == (model.User{}){
	  	  fmt.Println("结构体为空")
	  	 str :=  errmsg.GetCodeMsg(1002)
	  	 return str
	  }

	  if userModel != (model.User{}){
            if userModel.UserPaword != user.UserPaword {
            	str := errmsg.GetCodeMsg(1003)
            	return str
			}

			if user.UserName == userModel.UserName && userModel.UserPaword ==user.UserPaword{
				return "登陆成功"
			}
	  }


	  return ""
}