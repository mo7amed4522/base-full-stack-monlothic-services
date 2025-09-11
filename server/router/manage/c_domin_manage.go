package manage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CDominManageRouter struct{}

// InitCDominManageRouter 初始化 cDominManage表 路由信息
func (s *CDominManageRouter) InitCDominManageRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	cDominManageRouter := Router.Group("cDominManage").Use(middleware.OperationRecord())
	cDominManageRouterWithoutRecord := Router.Group("cDominManage")
	cDominManageRouterWithoutAuth := PublicRouter.Group("cDominManage")
	{
		cDominManageRouter.POST("createCDominManage", cDominManageApi.CreateCDominManage)             // 新建cDominManage表
		cDominManageRouter.DELETE("deleteCDominManage", cDominManageApi.DeleteCDominManage)           // 删除cDominManage表
		cDominManageRouter.DELETE("deleteCDominManageByIds", cDominManageApi.DeleteCDominManageByIds) // 批量删除cDominManage表
		cDominManageRouter.PUT("updateCDominManage", cDominManageApi.UpdateCDominManage)              // 更新cDominManage表
		cDominManageRouter.PUT("visibleCDominManage", cDominManageApi.VisibleCDominManage)            // 更新cDominManage表
	}
	{
		cDominManageRouterWithoutRecord.GET("findCDominManage", cDominManageApi.FindCDominManage)       // 根据ID获取cDominManage表
		cDominManageRouterWithoutRecord.GET("getCDominManageList", cDominManageApi.GetCDominManageList) // 获取cDominManage表列表
	}
	{
		cDominManageRouterWithoutAuth.GET("getCDominManagePublic", cDominManageApi.GetCDominManagePublic) // cDominManage表开放接口
	}
}
