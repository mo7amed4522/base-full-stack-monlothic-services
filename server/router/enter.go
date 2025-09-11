package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/bxapi"
	"github.com/flipped-aurora/gin-vue-admin/server/router/manage"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System system.RouterGroup
	Manage manage.RouterGroup
	BxApi  bxapi.RouterGroup
}
