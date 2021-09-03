package library

import "zhangyudevops.com/services"

type ApiGroup struct {
	BooksApi
	UserApi
}

var LibService = services.ServiceGroupApp.BooksServiceGroup.ManageBooksService