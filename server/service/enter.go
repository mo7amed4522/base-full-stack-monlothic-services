package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/bxapi"
	"github.com/flipped-aurora/gin-vue-admin/server/service/manage"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	ManageServiceGroup manage.ServiceGroup
	BxapiServiceGroup  bxapi.ServiceGroup
}
