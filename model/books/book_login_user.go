package books

import "zhangyudevops.com/global"

type BookLoginUser struct {
	global.GVA_MODEL
	UserName    string `json:"userName" gorm:"comment:用户登录名"`
	Password    string `json:"-"  gorm:"comment:用户登录密码"`
	PhoneNumber string    `json:"phoneNumber"  gorm:"comment:登录用户电话号码"`
}
