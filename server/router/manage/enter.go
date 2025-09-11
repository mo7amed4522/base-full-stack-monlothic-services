package manage

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type RouterGroup struct {
	CAppManageRouter
	CDominManageRouter
	CSubDominManageRouter
	CPageTemplateRouter
	CPromotionLinkRouter
	CBxMatchesRouter
	CBxAdvertisingRouter
}

var (
	CAppManageApi      = api.ApiGroupApp.ManageGroup.CAppManager
	cDominManageApi    = api.ApiGroupApp.ManageApiGroup.CDominManageApi
	cSubDominManageApi = api.ApiGroupApp.ManageApiGroup.CSubDominManageApi
	cPageTemplateApi   = api.ApiGroupApp.ManageApiGroup.CPageTemplateApi
	cPromotionLinkApi  = api.ApiGroupApp.ManageApiGroup.CPromotionLinkApi
	cBxMatchesApi      = api.ApiGroupApp.ManageApiGroup.CBxMatchesApi
	cBxAdvertisingApi  = api.ApiGroupApp.ManageApiGroup.CBxAdvertisingApi
)
