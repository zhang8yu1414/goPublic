package response

import "zhangyudevops.com/model/books"

type LoginResponse struct {
	User      books.BookLoginUser `json:"user"`
	Token     string              `json:"token"`
	ExpiresAt int64               `json:"expiresAt"`
}
