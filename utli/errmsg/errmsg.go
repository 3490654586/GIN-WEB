package errmsg

const (
	SUCCSE = 200
	ERROR  = 500

	//1000...用户错误
	ERROR_USERNAME_NOT_EXIST = 1001
	ERROR_USERNAME_EXISTHENCE = 1002
	ERROR_USERPASORD_WRONG   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
)

var codeMsg = map[int]string{
      SUCCSE:                    "OK",
      ERROR:                     "ERROR",
      ERROR_USERNAME_NOT_EXIST:  "用户名已经存在",
	  ERROR_USERNAME_EXISTHENCE: "用户名不存在",
       ERROR_USERPASORD_WRONG:    "密码错误",
		ERROR_TOKEN_EXIST:      "TOKEN不存在,请重新登陆",
		ERROR_TOKEN_RUNTIME:    "TOKEN已过期,请重新登陆",
		ERROR_TOKEN_WRONG:      "TOKEN不正确,请重新登陆",
		ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误,请重新登陆",
}

func GetCodeMsg(code int)string{
   return codeMsg[code]
}