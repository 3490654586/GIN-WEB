package model



type User struct {
	Id int `xorm:"pk autoincr" json:"id"`
	UserName string `xorm:"varchar(20) NOT NULL" json:"user_name" form:"user_name"`
	UserPaword string `xorm:"varchar(20) NOT NULL" json:"user_paword" form:"user_paword"`
}
