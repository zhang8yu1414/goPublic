package books

import (
	"time"
)

type Students struct {
	CreateAt     time.Time
	UpdateAt     time.Time
	DeleteAt     *time.Time `sql:"index"`
	Name         string     `json:"name" gorm:"comment:学生姓名"`
	StudentId    string     `json:"studentId" gorm:"comment:学生ID;not null;unique;primary_key"`
	Gender       string     `json:"gender" gorm:"comment:学生性别"`
	Age          string     `json:"age" gorm:"comment:学生年龄"`
	Phone        string     `json:"phone" gorm:"comment:学生电话"`
	CreationTime time.Time  `json:"creationTime" gorm:"comment:注册时间"`
}
