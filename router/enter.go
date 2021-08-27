package router

import "zhangyudevops.com/router/example"

type RouterGroup struct {
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)