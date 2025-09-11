package manage

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	CAppManager
	CDominManageApi
	CSubDominManageApi
	CPageTemplateApi
	CPromotionLinkApi
	CBxMatchesApi
	CBxAdvertisingApi
}

var (
	cAppManageService      = service.ServiceGroupApp.ManageServiceGroup.CAppManageService
	cDominManageService    = service.ServiceGroupApp.ManageServiceGroup.CDominManageService
	cSubDominManageService = service.ServiceGroupApp.ManageServiceGroup.CSubDominManageService
	cPageTemplateService   = service.ServiceGroupApp.ManageServiceGroup.CPageTemplateService
	cPromotionLinkService  = service.ServiceGroupApp.ManageServiceGroup.CPromotionLinkService
	cBxMatchesService      = service.ServiceGroupApp.ManageServiceGroup.CBxMatchesService
	cBxAdvertisingService  = service.ServiceGroupApp.ManageServiceGroup.CBxAdvertisingService
)
