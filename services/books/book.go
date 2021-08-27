package books

import (
	"zhangyudevops.com/global"
	"zhangyudevops.com/model/books"
)

type ManageBooksService struct {

}

func (m *ManageBooksService) CreateLoginUser(userName string, password string, phoneNumber string) (err error) {
	var loginUser books.BookLoginUser
	loginUser.UserName = userName
	loginUser.Password = password
	loginUser.PhoneNumber = phoneNumber

	err = global.GVA_DB.Create(&loginUser).Error
	return err
}

