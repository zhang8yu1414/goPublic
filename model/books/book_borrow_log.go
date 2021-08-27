package books

import (
	"github.com/satori/go.uuid"
	"time"
	"zhangyudevops.com/global"
)

type BookBorrowLog struct {
	global.GVA_MODEL
	ReaderNumber     uuid.UUID     `json:"readerNumber" gorm:"comment:学生借阅编号"`
	ReaderName       string    `json:"readName" gorm:"comment:学生姓名"`
	Students         Students  `json:"students" gorm:"comment:学生信息;foreignKey:StudentId;references:StudentId"`
	StudentId        string    `json:"studentId" gorm:"comment:学生ID"`
	BookName         string    `json:"bookName" gorm:"comment:图书书籍名称"`
	ReaderBorrowDate time.Time `json:"readerBorrowDate" gorm:"comment:学生借阅图书时间"`
}
