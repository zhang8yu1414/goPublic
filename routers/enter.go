package routers

import (
	"zhangyudevops.com/routers/example"
	"zhangyudevops.com/routers/system"
)

type RouterGroup struct {
	Example example.RouterGroup
	System  system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
