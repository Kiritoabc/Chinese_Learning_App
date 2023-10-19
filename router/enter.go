package router

import "Chinese_Learning_App/router/system"

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
