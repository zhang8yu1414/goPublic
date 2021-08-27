package books

import "time"

type PublishInfo struct {
	CreateAt    time.Time
	UpdateAt    time.Time
	DeleteAt    *time.Time `sql:"index"`
	PublishName string     `json:"publishName" gorm:"comment:图书出版社名称"`
	PublishId   string     `json:"publishId" gorm:"comment:图书出版社ID;not null;unique;primary_key"`
}
