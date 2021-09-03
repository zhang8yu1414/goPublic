package routers

import (
	"zhangyudevops.com/routers/library"
	"zhangyudevops.com/routers/system"
)

type RouterGroup struct {
	Example library.RouterGroup
	System  system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
