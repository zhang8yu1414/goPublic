package books

import (
	"time"
	"zhangyudevops.com/global"
)

type Students struct {
	global.GVA_MODEL
	Name         string     `json:"name" gorm:"comment:学生姓名"`
	Gender       string     `json:"gender" gorm:"comment:学生性别"`
	Age          string     `json:"age" gorm:"comment:学生年龄"`
	Phone        string     `json:"phone" gorm:"comment:学生电话"`
	CreationTime time.Time  `json:"creationTime" gorm:"comment:注册时间"`
}
