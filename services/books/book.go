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

// SelectLoginUser 用户登录
func (m ManageBooksService) SelectLoginUser(username string, password string) (userInfo *books.BookLoginUser, err error) {
	var user books.BookLoginUser
	err = global.GVA_DB.Where("user_name = ? AND password = ?", username, password).First(&user).Error
	return &user, err
}

// DeleteLoginUser 删除用户
func (m ManageBooksService) DeleteLoginUser(username string) (userInfo *books.BookLoginUser, err error) {
	var user books.BookLoginUser
	err = global.GVA_DB.Where("user_name = ?", username).Delete(&user).Error
	return &user, err
}
