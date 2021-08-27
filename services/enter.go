package services

import "zhangyudevops.com/services/books"

type ServiceGroup struct {
	BooksServiceGroup books.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)