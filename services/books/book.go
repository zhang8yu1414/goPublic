package books

import (
	"zhangyudevops.com/global"
	"zhangyudevops.com/model/books"
)

type ManageBooksService struct {

}

// CreateLoginUser 存入数据库
func (m *ManageBooksService) CreateLoginUser(user books.BookLoginUser) (err error) {

	err = global.GVA_DB.Create(&user).Error
	return err
}

func (m ManageBooksService) SelectLoginUser()  {

}

