package library

import "zhangyudevops.com/services"

type ApiGroup struct {
	BooksApi
}

var LibService = services.ServiceGroupApp.BooksServiceGroup.ManageBooksService