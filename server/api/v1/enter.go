package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/manage"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	ManageGroup    manage.ApiGroup
	ManageApiGroup manage.ApiGroup
}
