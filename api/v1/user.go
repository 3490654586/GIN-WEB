package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wep-app/model"
	"wep-app/pkg/jwt"
	"wep-app/server"
	"wep-app/utli/errmsg"
)

func SupnIn(c *gin.Context)  {
	 var user model.User
	 
	err := c.ShouldBind(&user)
	fmt.Println("err=",err)
	fmt.Println("user=",user)
	if err != nil {
		c.JSON(200,gin.H{
			"msg": "解析参数失败",
		})
		return
	}

	 str := server.CheckUser(user)
    fmt.Println(str)
	switch str{
	case errmsg.GetCodeMsg(1001):
		c.JSON(200,gin.H{
			"msg":str,
		})
	case errmsg.GetCodeMsg(1002):
		c.JSON(200,gin.H{
			"msg":str,
		})
	case errmsg.GetCodeMsg(1003):
		c.JSON(200,gin.H{
			"msg":str,
		})
	default:
		token,err := jwt.SetToken(user.UserName)
		if err != nil {
			return
		}
		c.JSON(200,gin.H{
			"msg":str,
			"token":token,
		})
	}
}
