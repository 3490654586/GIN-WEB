package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"wep-app/utli/errmsg"
)

const Userkey  = "用户名"

type MyClaims struct {
	//自定义字段
	username string
	//官方字段
	jwt.StandardClaims
}

//加上我们的加密数据
var mySecret=[]byte("我是帅哥")

//生成touken
func SetToken (username string)(string,error){
	//初始化生成token的内容
     myclaims := &MyClaims{
		 username:      username,
		 StandardClaims: jwt.StandardClaims{
		 	//过期时间
            ExpiresAt:time.Now().Add(7 * 24 * time.Hour).Unix(),
            //发布人
			 Issuer:"张宪冲",
		 },
	 }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,myclaims)
return token.SignedString(mySecret)
}

//验证token
func GetToken(tokenString string)(*MyClaims,string){
	//解析到结构体
	var myclaims = new(MyClaims)

	token,err := jwt.ParseWithClaims(tokenString,myclaims, func(token *jwt.Token) (i interface{}, e error) {
		return mySecret,nil
	})

	if err != nil {
		error := fmt.Sprintf("%s", err)
		return nil,error
	}

	//Valid代表token是否可用返回布尔值
	if token.Valid{
		return myclaims,""
	}

	return nil,errmsg.GetCodeMsg(1004)
}

//JWT中间件
func JWTAuthMiddleware() func( c *gin.Context) {
    return func(c *gin.Context) {

		//从请求头获取token
		authHhande := c.Request.Header.Get("Authorization")
		if authHhande == ""{
			c.JSON(200,gin.H{
				"msg":errmsg.GetCodeMsg(1004),
			})
			c.Abort()
			return
		}

		//因为token前面加了Bearer字符按空格分割在进行校验
		parts := strings.SplitN(authHhande," ",2)
		fmt.Println("psrts[0]",parts[0])
		fmt.Println("psrts[0]",parts[1])
		if !(len(parts) == 2 && parts[0]=="Bearer"){
			c.JSON(200,gin.H{
				"msg":errmsg.GetCodeMsg(1006),
			})
			c.Abort()
			return
		}

		//parts[1]真实token
		token,str := GetToken(parts[1])
		if str != "" {
			c.JSON(200,gin.H{
				"msg":errmsg.GetCodeMsg(1006),
			})
			c.Abort()
			return
		}
		//将用户名字保存到c的上下文
		c.Set(Userkey,token.username)
		// 后续的处理请求的函数中 可以用过c.Get(Userkey) 来获取当前请求的用户信息
		c.Next()
	}
}