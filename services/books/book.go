package books

import (
	"gorm.io/gorm"
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

// SelectLoginUser 查询用户登录数据
func (m ManageBooksService) SelectLoginUser(username string, password string) *gorm.DB {
	var user books.BookLoginUser
	result := global.GVA_DB.Where(&books.BookLoginUser{
		UserName: username,
		Password: password,
	}).First(&user)
	return result
}

func (m ManageBooksService) DeleteLoginUser(username string) *gorm.DB {
	var user books.BookLoginUser
	result := global.GVA_DB.Where("user_name = ?", username).Delete(&user)
	return result
}
