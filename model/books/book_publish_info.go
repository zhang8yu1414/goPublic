package books

import "zhangyudevops.com/global"

type PublishInfo struct {
	global.GVA_MODEL
	PublishName string     `json:"publishName" gorm:"comment:图书出版社名称"`
	//PublishId   string     `json:"publishId" gorm:"comment:图书出版社ID;not null;unique;primary_key"`
}
