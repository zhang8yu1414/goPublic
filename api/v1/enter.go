package v1

import "zhangyudevops.com/api/v1/example"

type ApiGroup struct {
	TestApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)