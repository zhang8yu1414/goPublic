package v1

import "zhangyudevops.com/api/v1/library"

type ApiGroup struct {
	TestApiGroup library.ApiGroup
}

var ApiGroupApp = new(ApiGroup)