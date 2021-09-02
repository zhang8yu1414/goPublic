package request

type Login struct {
	Username  string `json:"userName"`  // 用户名
	Password  string `json:"password"`  // 密码
}

type Register struct {
	Username  string `json:"userName"`  // 用户名
	Password  string `json:"password"`  // 密码
	PhoneNumber string `json:"phoneNumber"` // 电话号码
}
