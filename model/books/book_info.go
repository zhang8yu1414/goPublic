package books

import "zhangyudevops.com/global"

type BookInfo struct {
	global.GVA_MODEL
	Name        string      `json:"name" gorm:"comment:图书名称"`
	Author      string      `json:"author" gorm:"comment:图书作者"`
	Publish     PublishInfo `json:"publish" gorm:"comment:图书出版社;foreignKey:ID;references:ID"`
	PublishId   string      `json:"publishId" gorm:"comment:图书出版社ID"`
	Type        string      `json:"type" gorm:"comment:图书类型"`
	Count       int         `json:"count" gorm:"comment:图书总量"`
	RemainCount int         `json:"remainCount" gorm:"comment:图书库存数量"`
	Price       float64     `json:"price" gorm:"comment:图书价格"`
}
