package manage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CPromotionLinkRouter struct {}

// InitCPromotionLinkRouter 初始化 cPromotionLink表 路由信息
func (s *CPromotionLinkRouter) InitCPromotionLinkRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	cPromotionLinkRouter := Router.Group("cPromotionLink").Use(middleware.OperationRecord())
	cPromotionLinkRouterWithoutRecord := Router.Group("cPromotionLink")
	cPromotionLinkRouterWithoutAuth := PublicRouter.Group("cPromotionLink")
	{
		cPromotionLinkRouter.POST("createCPromotionLink", cPromotionLinkApi.CreateCPromotionLink)   // 新建cPromotionLink表
		cPromotionLinkRouter.DELETE("deleteCPromotionLink", cPromotionLinkApi.DeleteCPromotionLink) // 删除cPromotionLink表
		cPromotionLinkRouter.DELETE("deleteCPromotionLinkByIds", cPromotionLinkApi.DeleteCPromotionLinkByIds) // 批量删除cPromotionLink表
		cPromotionLinkRouter.PUT("updateCPromotionLink", cPromotionLinkApi.UpdateCPromotionLink)    // 更新cPromotionLink表
	}
	{
		cPromotionLinkRouterWithoutRecord.GET("findCPromotionLink", cPromotionLinkApi.FindCPromotionLink)        // 根据ID获取cPromotionLink表
		cPromotionLinkRouterWithoutRecord.GET("getCPromotionLinkList", cPromotionLinkApi.GetCPromotionLinkList)  // 获取cPromotionLink表列表
	}
	{
	    cPromotionLinkRouterWithoutAuth.GET("getCPromotionLinkPublic", cPromotionLinkApi.GetCPromotionLinkPublic)  // cPromotionLink表开放接口
	}
}
