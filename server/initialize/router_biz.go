package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	bxPrivateGroup := routers[2]
	holder(publicGroup, privateGroup, bxPrivateGroup)
	{
	}
	{
	}
	{
		manageRouter := router.RouterGroupApp.Manage
		manageRouter.InitAppManageRouter(privateGroup, publicGroup)
		manageRouter.InitCDominManageRouter(privateGroup, publicGroup)
		manageRouter.InitCSubDominManageRouter(privateGroup, publicGroup)
		manageRouter.InitCPageTemplateRouter(privateGroup, publicGroup)
		manageRouter.InitCPromotionLinkRouter(privateGroup, publicGroup)
		manageRouter.InitBxMatchesRouter(privateGroup, publicGroup)
		manageRouter.InitBxAdvertisingRouter(privateGroup, publicGroup)
	}
	{
		bxApiRouter := router.RouterGroupApp.BxApi
		bxApiRouter.InitBxUserRouter(bxPrivateGroup, publicGroup)
		bxApiRouter.InitBxMatchesRouter(bxPrivateGroup, publicGroup) // 使用 bxPrivateGroup 进行 bxapi 认证
		bxApiRouter.InitBxVotesRouter(bxPrivateGroup, publicGroup)
		bxApiRouter.InitBxTipsRouter(bxPrivateGroup, publicGroup)
		bxApiRouter.InitBxRewardRouter(bxPrivateGroup, publicGroup)
		bxApiRouter.InitGoogleAuthRouter(bxPrivateGroup, publicGroup) // Google认证路由
		bxApiRouter.InitBxConfigRouter(bxPrivateGroup, publicGroup)   // 配置路由
	}
	{
	}
}
